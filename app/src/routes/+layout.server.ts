import {generateAuthUrl, hydrateAuth, isSignedIn} from 'svelte-google-auth/server';
import type { LayoutServerLoad } from './$types.js';
import {redirect} from "@sveltejs/kit";

const SCOPES = ['openid', 'profile', 'email','https://mail.google.com/'];

export const load: LayoutServerLoad = ({ locals, url }) => {
    if (!isSignedIn(locals)) {
        throw redirect(302, generateAuthUrl(locals, url, SCOPES, url.pathname));
    }

    // By calling hydrateAuth, certain variables from locals are parsed to the client
    // allowing the client to access the user information and the client_id for login
    return { ...hydrateAuth(locals) };
};