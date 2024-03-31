import { apiUrl } from '$lib/api.js';

type Deck = {
    id: string;
    title: string;
    cards: [string, string][];
    createdAt: string;
    updatedAt: string;
};

export const load = async ({ fetch, params }) => {
    let r = await fetch(apiUrl + '/deck/' + params.id, {
        method: 'get'
    });
    let data: Deck = await r.json();

    return { deck: data };
};
