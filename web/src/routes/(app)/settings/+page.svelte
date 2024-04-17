<script lang="ts">
	import ThemeDropdown from '$lib/components/ThemeDropdown.svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	const logOut = () => {
		document.cookie = `sessionID=; max-age=0; path=/; secure=true;`;
		localStorage.removeItem('username');
		window.location.href = '/';
	};
</script>

<div class="flex h-full w-full justify-center">
	<div class="flex w-full max-w-[40rem] flex-col gap-4 pt-6">
		<Button on:click={logOut} variant="destructive" data-sveltekit-reload>Log out</Button>
	</div>
</div>
