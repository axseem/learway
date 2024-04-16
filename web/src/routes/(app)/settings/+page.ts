import { goto } from '$app/navigation';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ parent }) => {
	const { authorized } = await parent();
	if (authorized == false) {
		goto('/login');
	}
};
