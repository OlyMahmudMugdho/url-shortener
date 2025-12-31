import { inject } from '@angular/core';
import { Router, CanActivateFn } from '@angular/router';
import { AuthService } from './auth.service';

export const authGuard: CanActivateFn = (route, state) => {
    const authService = inject(AuthService);
    const router = inject(Router);

    if (authService.isAuthenticated) {
        return true;
    } else {
        // Ideally, we might want to check if the session is valid by making an API call 
        // if the state is lost on refresh, but for now we rely on the client-side state.
        // A better approach for persistence is to check a cookie/token or have an /me endpoint.
        // Since we are using httpOnly cookies, we can try to hit a protected endpoint or check a local flag.
        // Given the current setup, if 'isAuthenticated' is false, redirect to login.
        return router.createUrlTree(['/login']);
    }
};
