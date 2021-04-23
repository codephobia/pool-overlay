import {
  trigger,
  animate,
  style,
  group,
  query,
  transition,
  sequence,
  animation,
  useAnimation
} from '@angular/animations';

const scoreboardAnimation = animation([
  sequence([
    group([
      query(':enter .player-wrapper.one', style({ width: '22px' }), { optional: true }),
      query(':enter .player-wrapper.two', style({ width: '22px' }), { optional: true }),
      query(':enter .scores-wrapper', [
        style({ height: '0px' }),
        animate('0.5s ease-in-out', style({ height: '*' })),
      ], { optional: true }),
      query(':enter .game', [
        style({ transform: 'translateY(-100%)' }),
        animate('0.5s ease-in-out', style({ transform: 'translateY(0)' })),
      ], { optional: true }),
    ]),
    group([
      query(':enter .player-wrapper.one', [
        style({ width: '22px' }),
        animate('0.5s ease-in-out', style({ width: '*' })),
      ], { optional: true }),
      query(':enter .player-wrapper.two', [
        style({ width: '22px' }),
        animate('0.5s ease-in-out', style({ width: '*' })),
      ], { optional: true }),
      query(':leave .player-wrapper.one', [
        style({ width: '*' }),
        animate('0.5s ease-in-out', style({ width: '22px' })),
      ], { optional: true }),
      query(':leave .player-wrapper.two', [
        style({ width: '*' }),
        animate('0.5s ease-in-out', style({ width: '22px' })),
      ], { optional: true }),
    ]),
    group([
      query(':leave .scores-wrapper', [
        style({ height: '*' }),
        animate('0.5s ease-in-out', style({ height: '0px' })),
      ], { optional: true }),
      query(':leave .game', [
        style({ transform: 'translateY(0)' }),
        animate('0.5s ease-in-out', style({ transform: 'translateY(-100%)' })),
      ], { optional: true }),
    ]),
  ]),
]);

export const scoreboardTransition = trigger('scoreboardTransition', [
  transition('void => *', []),
  transition('* => *', [
    useAnimation(scoreboardAnimation)
  ]),
]);
