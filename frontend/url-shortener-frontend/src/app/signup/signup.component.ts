import { Component } from '@angular/core';
import { ButtonDirective } from "primeng/button";
import { FormsModule } from "@angular/forms";
import { IconFieldModule } from "primeng/iconfield";
import { InputIconModule } from "primeng/inputicon";
import { InputTextModule } from "primeng/inputtext";
import { Ripple } from "primeng/ripple";
import { NgStyle } from "@angular/common";
import { Router, RouterLink } from "@angular/router";
import { AuthService } from "../services/auth/auth.service";

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
  username: any;
  password: any;
  firstName: any;
  lastName: any;
  email: any;

  constructor(private authService: AuthService, private router: Router) { }

  handleSubmit() {
    const user = {
      username: this.username,
      password: this.password,
      first_name: this.firstName,
      last_name: this.lastName,
      email: this.email
    };

    this.authService.register(user).subscribe({
      next: (response) => {
        alert("Registration successful!");
        this.router.navigate(['/login']);
      },
      error: (error) => {
        console.error("Registration failed", error);
        alert("Registration failed: " + (error.error?.message || "Unknown error"));
      }
    });
  }
}
