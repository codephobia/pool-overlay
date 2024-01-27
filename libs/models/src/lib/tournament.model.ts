import { TournamentMatch } from './tournament-match.model';

export interface Tournament {
    id: number;
    name: string;
    url: string;
    tournament_type: string;
    rounds: number;
    matches: TournamentMatch[] | null;
    created_at: string;
    completed_at: string | null;
}
