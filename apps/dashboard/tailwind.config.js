module.exports = {
    purge: {
        enabled: false,
        content: ['./apps/dashboard/src/**/*.{html,ts}'],
    },
    theme: {
        extend: {
            colors: {
                'grad-start': 'var(--color-sad-bg-grad-start)',
                'grad-stop': 'var(--color-sad-bg-grad-stop)',
            },
            textColor: {
                sad: {
                    icon: 'var(--color-sad-text)',
                },
            },
            backgroundColor: {
                sad: {
                    background: 'var(--color-sad-background)',
                    active: 'var(--color-sad-background-active)',
                    input: 'var(--color-sad-input-background)',
                    'section-title':
                        'var(--color-sad-section-title-background)',
                    'table-odd': 'var(--color-sad-bg-table-odd)',
                    'table-even': 'var(--color-sad-bg-table-even)',
                    success: 'var(--color-sad-success)',
                    error: 'var(--color-sad-error)',
                    challonge: 'var(--color-sad-challonge)',
                    'player-one': 'var(--color-sad-player-one-background)',
                    'player-two': 'var(--color-sad-player-two-background)',
                    'player-score': 'var(--color-sad-player-score-background)',
                },
            },
            borderColor: {
                sad: {
                    dark: 'var(--color-sad-border)',
                    input: 'var(--color-sad-input-background)',
                    active: 'var(--color-sad-border-active)',
                },
            },
            width: {
                '42px': '42px',
            },
            height: {
                '42px': '42px',
                '66px': '66px',
            },
            maxWidth: {
                '400px': '400px',
            },
            borderRadius: {
                DEFAULT: '9px',
            },
            padding: {
                '14px': '14px',
                '50px': '50px',
            },
        },
    },
    variants: {},
    plugins: [],
};
