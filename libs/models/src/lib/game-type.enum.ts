export enum GameType {
    EightBall,
    NineBall,
    TenBall,
}

const GameTypeName = [
    '8 Ball',
    '9 Ball',
    '10 Ball',
];

export function gameTypeToString(gameType: GameType): string {
    return GameTypeName[gameType];
}
