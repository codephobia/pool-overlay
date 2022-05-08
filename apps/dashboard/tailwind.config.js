module.exports = {
    purge: {
        enabled: false,
        content: ['./apps/dashboard/src/**/*.{html,ts}'],
    },
    theme: {
        extend: {
            textColor: {
                nav: {
                    icon: 'var(--color-nav-icon)',
                },
            },
            backgroundColor: {
                nav: {
                    background: 'var(--color-nav-background)',
                    'background-active': 'var(--color-nav-background-active)',
                },
            },
            borderColor: {
                nav: {
                    dark: 'var(--color-nav-border)',
                },
            },
        },
    },
    variants: {},
    plugins: [],
};
