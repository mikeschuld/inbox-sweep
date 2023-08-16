import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types.js';
import {isSignedIn} from "svelte-google-auth/server";

export const load: PageServerLoad = ({ locals }) => {
    if (!isSignedIn(locals)) throw error(403, 'Not signed in');
};