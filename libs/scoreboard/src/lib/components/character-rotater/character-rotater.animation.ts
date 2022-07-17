import { trigger, style, transition, query, stagger, animate, animation, group, useAnimation } from '@angular/animations';

const characterRotaterAnimation = animation([
    group([
        query('.characters:enter .character', style({
            transform: 'rotateX(-90deg)',
            opacity: 0,
        }), { optional: true }),
        query('.characters:enter .character', [
            stagger(100, [
                animate('0.38s 0.32s ease-in-out', style({
                    transform: 'rotateX(0deg)',
                    opacity: 1,
                })),
            ])
        ], { optional: true }),
        query('.characters:leave .character', style({
            transform: 'rotateX(0deg)',
            opacity: 1
        }), { optional: true }),
        query('.characters:leave .character', [
            stagger(100, [
                animate('0.32s cubic-bezier(0.55, 0.055, 0.675, 0.19)', style({
                    transform: 'rotateX(90deg)',
                    opacity: 1,
                })),
            ])
        ], { optional: true }),
    ])
]);

export const characterRotater = trigger('characterRotater', [
    transition('void => *', []),
    transition('* => void', []),
    transition('* => *', [
        useAnimation(characterRotaterAnimation)
    ]),
]);
