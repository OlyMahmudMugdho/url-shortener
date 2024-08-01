import { Component } from '@angular/core';
import {ButtonDirective} from "primeng/button";
import {FormsModule} from "@angular/forms";
import {IconFieldModule} from "primeng/iconfield";
import {InputIconModule} from "primeng/inputicon";
import {InputTextModule} from "primeng/inputtext";
import {Ripple} from "primeng/ripple";
import {NgStyle} from "@angular/common";
import {RouterLink} from "@angular/router";

@Component({
  selector: 'app-signup',
  standalone: true,
  imports: [
    ButtonDirective,
    FormsModule,
    IconFieldModule,
    InputIconModule,
    InputTextModule,
    Ripple,
    NgStyle,
    RouterLink
  ],
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.css'
})
export class SignupComponent {
  username:any;
  password:any;
  firstName: any;
  lastName: any;
  email: any;

  handleSubmit(){
    alert(this.username + " : " + this.password)
  }
}
