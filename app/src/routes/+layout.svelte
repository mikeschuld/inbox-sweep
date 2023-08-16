<script lang="ts">
    import { invalidateAll } from '$app/navigation';
    import type { PageData } from './$types.js';
    import {initialize, signIn, signOut} from "svelte-google-auth/client";

    export let data: PageData;
    initialize(data, invalidateAll);

    console.log('data', data);
</script>

<header>
    {#if data.auth.user}
        <p>
            <img src={data.auth.user.picture} width={36} referrerpolicy="no-referrer" alt="profile" />
        </p>
        <p>
            Signed in as {data.auth.user.name} ({data.auth.user.email})
        </p>
        <p>
            <button on:click={() => signOut()}>Sign out</button>
        </p>
    {:else}
        <p>
            <button
                    on:click={() =>
					signIn([
						'openid',
						'profile',
						'email',
						'https://mail.google.com/'
					])}>Sign in</button
            >
        </p>
    {/if}
</header>

<slot />

<svelte:head>
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous" />
    <link href="https://fonts.googleapis.com/css2?family=Roboto&display=swap" rel="stylesheet" />
</svelte:head>

<style>
    :global(body) {
        font-family: 'Roboto', sans-serif;
        max-width: 800px;
        margin: auto;
        padding: 1rem;
    }

    header {
        border-bottom: 1px solid #ccc;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    header p {
        display: flex;
        align-items: center;
    }
</style>