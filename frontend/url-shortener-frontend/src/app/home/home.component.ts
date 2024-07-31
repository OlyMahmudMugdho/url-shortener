import { Component } from '@angular/core';
import {Button} from "primeng/button";
import {ImageModule} from "primeng/image";
import {NgStyle} from "@angular/common";

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [
    Button,
    ImageModule,
    NgStyle
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {

}
