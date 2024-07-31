import { Component } from '@angular/core';
import {NgStyle} from "@angular/common";

@Component({
  selector: 'app-footer',
  standalone: true,
  imports: [
    NgStyle
  ],
  templateUrl: './footer.component.html',
  styleUrl: './footer.component.css'
})
export class FooterComponent {

}
