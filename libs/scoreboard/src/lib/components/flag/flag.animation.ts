import { trigger, style, transition, query, animate, animation, group, useAnimation } from '@angular/animations';

const flagAnimation = animation([
    group([
        query('img:enter', style({
            transform: 'translateX({{offsetEnter}})',
        }), { optional: true }),
        query('img:enter', [
            animate('.76s ease-in-out', style({
                transform: 'translateX(0)',
            })),
        ], { optional: true }),
        query('img:leave', style({
            transform: 'translateX(0)',
        }), { optional: true }),
        query('img:leave', [
            animate('.76s ease-in-out', style({
                transform: 'translateX({{offsetLeave}})',
            })),
        ], { optional: true }),
    ])
]);

export const flagTrigger = trigger('flagTrigger', [
    transition('void => *', []),
    transition('* => void', []),
    transition('* => *', [
        useAnimation(flagAnimation)
    ]),
]);
