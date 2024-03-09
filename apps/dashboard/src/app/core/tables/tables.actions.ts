import { createActionGroup, emptyProps, props } from '@ngrx/store';

export const TablesActions = createActionGroup({
    source: 'Tables',
    events: {
        'Get Count': emptyProps(),
        'Get Count Success': props<{ count: number }>(),
        'Get Count Error': props<{ error: string }>(),
        'Set Count': props<{ count: number }>(),
    },
});
