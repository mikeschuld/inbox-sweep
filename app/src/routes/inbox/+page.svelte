<script lang="ts">
    import { getGapiClient } from "svelte-google-auth/client";
    import type { PageData } from './$types.js';
    export let data: PageData;

    let groupedEmails: Array<{ sender: string, count: number }> = [];

    async function retrieveAllEmails(client, pageToken = null) {
        const response = await client.gmail.users.messages.list({
            'userId': 'me',
            'pageToken': pageToken,
            'maxResults': 500,
            'fields': 'nextPageToken,messages/id,messages/labelIds'
        });

        const grouped = {};

        response.result.messages.forEach(message => {
            const senderHeader = message.payload.headers.find(header => header.name === 'From');
            if (senderHeader) {
                const sender = senderHeader.value;
                grouped[sender] = (grouped[sender] || 0) + 1;
            }
        });

        groupedEmails = Object.entries(grouped)
            .map(([sender, count]) => ({ sender, count }))
            .sort((a, b) => b.count - a.count);

        if (response.result.nextPageToken) {
            await retrieveAllEmails(client, response.result.nextPageToken);
        }
    }

    if (data.auth.user) {
        getGapiClient({
            discoveryDocs: ['https://www.googleapis.com/discovery/v1/apis/gmail/v1/rest'],
        })
            .then(client => {
                retrieveAllEmails(client);
            });
    }
</script>

<table>
    <thead>
    <tr>
        <th>Sender</th>
        <th>Count</th>
    </tr>
    </thead>
    <tbody>
    {#each groupedEmails as { sender, count }}
        <tr>
            <td>{sender}</td>
            <td>{count}</td>
        </tr>
    {/each}
    </tbody>
</table>

{#if !data.auth.user}
    Not signed in
{/if}
