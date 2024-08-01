import { Component } from '@angular/core';
import {IconFieldModule} from "primeng/iconfield";
import {InputIconModule} from "primeng/inputicon";
import {InputTextModule} from "primeng/inputtext";
import {NgStyle} from "@angular/common";
import {FloatLabelModule} from "primeng/floatlabel";
import {Button, ButtonDirective} from "primeng/button";
import {Ripple} from "primeng/ripple";
import {FormsModule} from "@angular/forms";
import {RouterLink} from "@angular/router";

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    IconFieldModule,
    InputIconModule,
    InputTextModule,
    NgStyle,
    FloatLabelModule,
    Button,
    ButtonDirective,
    Ripple,
    FormsModule,
    RouterLink
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {
  username:any;
  password:any;

  handleSubmit(){
    alert(this.username + " : " + this.password)
  }
}
