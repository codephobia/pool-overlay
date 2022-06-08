module.exports = {
    purge: {
        enabled: false,
        content: ['./apps/dashboard/src/**/*.{html,ts}'],
    },
    theme: {
        extend: {
            textColor: {
                sad: {
                    icon: 'var(--color-sad-text)',
                },
            },
            backgroundColor: {
                sad: {
                    background: 'var(--color-sad-background)',
                    'background-active': 'var(--color-sad-background-active)',
                },
            },
            borderColor: {
                sad: {
                    dark: 'var(--color-sad-border)',
                },
            },
        },
    },
    variants: {},
    plugins: [],
};
