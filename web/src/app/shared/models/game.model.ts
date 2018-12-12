import {Board} from './board.model';

export class Game {
  id: number;
  email: string;
  board: Board;
  over: boolean;
  createdAt: Date;
  finishedAt: Date;
  elapsedTime: number;
}
