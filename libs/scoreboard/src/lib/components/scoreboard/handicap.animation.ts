import {
    trigger,
    animate,
    style,
    query,
    transition,
    animation,
    useAnimation
} from '@angular/animations';

const handicapAnimation = animation([
    query(':enter', [
        style({ transform: 'translateY(100%)' }),
        animate('0.5s ease-in-out', style({ transform: 'translateY(0)' })),
    ], { optional: true }),
    query(':leave', [
        style({ transform: 'translateY(0)' }),
        animate('0.5s ease-in-out', style({ transform: 'translateY(100%)' })),
    ], { optional: true }),
]);

export const handicapTransition = trigger('handicapTransition', [
    transition('* => *', [
        useAnimation(handicapAnimation)
    ]),
]);
