import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {Cell} from '../../../../shared/models/cell.model';

@Component({
  selector: 'app-cell',
  templateUrl: './cell.component.html',
  styleUrls: ['./cell.component.scss']
})
export class CellComponent implements OnInit {
  @Input() cell: Cell;
  @Output() openedCell: EventEmitter<Cell>;
  @Output() flaggedCell: EventEmitter<Cell>;

  constructor() {
    this.openedCell = new EventEmitter(null);
    this.flaggedCell = new EventEmitter(null);
  }

  ngOnInit() {
  }

  onClick() {
    this.openedCell.emit(this.cell);
  }

  onRightClick() {
    this.flaggedCell.emit(this.cell);
  }

  isRevealed(): string {
    if (this.cell.revealed) {
      return 'btn-outline';
    }

    return '';
  }
}
