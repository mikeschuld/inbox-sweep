import { SvelteGoogleAuthHook } from 'svelte-google-auth/server';
import type { Handle } from '@sveltejs/kit';

import {OAUTH_CLIENT_ID, OAUTH_CLIENT_SECRET} from "$env/static/private";

const auth = new SvelteGoogleAuthHook({
    client_id: OAUTH_CLIENT_ID,
    client_secret: OAUTH_CLIENT_SECRET
});

export const handle: Handle = async ({ event, resolve }) => {
    return auth.handleAuth({event, resolve});
};