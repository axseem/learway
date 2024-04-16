<script lang="ts">
	import { CirclePlus } from 'lucide-svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import { goto } from '$app/navigation';

	let { children, data } = $props();
	let authorized = $state(data.authorized);

	let searchValue: string = $state('');

	const search = (title: string) => {
		if (title.trim() != '') {
			goto('/search/' + title);
		}
	};
</script>

<div class="bg-background text-foreground flex h-screen ring-1">
	<aside class="flex h-full w-64 shrink-0 flex-col border-r border-t">
		<a class="border-b p-4 font-mono" href="/">learway</a>
		<form
			on:submit|preventDefault={() => {
				search(searchValue);
				searchValue = '';
			}}
			class="px-4 pt-4"
		>
			<Input bind:value={searchValue} placeholder="Search for knowledge" />
		</form>
		<div class="flex flex-col gap-2 p-4">
			<Button variant="ghost" href="/" class="justify-start">Recent</Button>
			<Button variant="ghost" href="/" class="justify-start">Popular</Button>
			<Button variant="ghost" href="/" class="justify-start">Latest</Button>
		</div>
		<div class="mt-auto flex flex-col gap-4 py-4">
			{#if authorized}
				<Button variant="secondary" href="/create" class="mx-4 gap-2">
					<CirclePlus size={16} />
					Create deck
				</Button>
				<Button
					variant="outline"
					href="/settings"
					class="mx-4 h-auto justify-start gap-2 rounded-lg border p-2"
				>
					<div class="bg-foreground h-8 w-8 rounded-full"></div>
					@{localStorage.getItem('username')}
				</Button>
			{:else}
				<Button href="/login" class="mx-4">Log In</Button>
			{/if}
		</div>
	</aside>
	<div class="h-screen w-full overflow-auto border-t p-4">
		{@render children()}
	</div>
</div>
