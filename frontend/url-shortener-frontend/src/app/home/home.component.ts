import { Component } from '@angular/core';
import {Button} from "primeng/button";
import {ImageModule} from "primeng/image";

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [
    Button,
    ImageModule
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {

}
