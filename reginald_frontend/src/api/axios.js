import axios from 'axios';

const api = axios.create({
    baseURL: '/api', // Correctly point to the prefixed gateway routes
    withCredentials: true, // Send cookies with every request
    headers: {
        'Content-Type': 'application/json',
    },
});

// Add a request interceptor to start timing
api.interceptors.request.use(
    (config) => {
        // Start timing the request
        config.metadata = { startTime: performance.now() };
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// Add response interceptor to calculate latency and handle errors
api.interceptors.response.use(
    async (response) => {
        // Calculate latency
        const endTime = performance.now();
        const startTime = response.config.metadata?.startTime || endTime;
        const durationMs = Math.round(endTime - startTime);

        // Log latency for debugging
        console.log(`[API] ${response.config.method?.toUpperCase()} ${response.config.url} - ${durationMs}ms`);

        // Send latency to backend for metrics (non-blocking)
        try {
            // We use fetch with credentials: 'include' to pass the token cookie
            // We need to ensure we prefix here as well if using fetch directly
            if (!response.config.url?.includes('/metrics/latency')) {
                // Fire and forget - don't await
                fetch('/api/metrics/latency', {
                    method: 'POST',
                    credentials: 'include',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        endpoint: response.config.url,
                        method: response.config.method?.toUpperCase(),
                        duration_ms: durationMs,
                        status_code: response.status
                    })
                }).catch(() => { }); // Silently ignore errors
            }
        } catch (e) {
            // Ignore errors when reporting latency
        }

        return response;
    },
    (error) => {
        // Calculate latency even for errors
        if (error.config?.metadata?.startTime) {
            const endTime = performance.now();
            const durationMs = Math.round(endTime - error.config.metadata.startTime);
            console.log(`[API ERROR] ${error.config.method?.toUpperCase()} ${error.config.url} - ${durationMs}ms - ${error.response?.status || 'Network Error'}`);

            // Report error latency too
            try {
                if (!error.config.url?.includes('/metrics/latency')) {
                    fetch('/api/metrics/latency', {
                        method: 'POST',
                        credentials: 'include',
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify({
                            endpoint: error.config.url,
                            method: error.config.method?.toUpperCase(),
                            duration_ms: durationMs,
                            status_code: error.response?.status || 500
                        })
                    }).catch(() => { });
                }
            } catch (e) { }
        }

        if (error.response?.status === 401) {
            console.warn('[AXIOS] 401 Unauthorized - Cookie may be expired or missing');
            localStorage.removeItem('user');
            if (!window.location.pathname.includes('/login')) {
                window.location.href = '/login';
            }
        }
        return Promise.reject(error);
    }
);

export default api;
