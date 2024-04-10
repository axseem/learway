import { apiUrl } from '$lib/api.js';

export const load = async ({ fetch, params }) => {
    let r = await fetch(apiUrl + '/user/' + params.id, {
        method: 'get'
    });
    let user = await r.json();

    return { user: params.id };
};
