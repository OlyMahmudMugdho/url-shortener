import { Component } from '@angular/core';
import { Button } from "primeng/button";
import { ImageModule } from "primeng/image";
import { NgStyle } from "@angular/common";
import { RouterLink } from "@angular/router";

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [
    Button,
    ImageModule,
    NgStyle,
    RouterLink
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent {

}
