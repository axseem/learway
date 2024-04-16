<script lang="ts">
	import { goto } from '$app/navigation';
	import Badge from '$lib/components/ui/badge/badge.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import { Search } from 'lucide-svelte';

	let { data } = $props();
	let decks = $state(data.decks);
	let searchValue: string = $state('');

	const search = (title: string) => {
		if (title.trim() != '') {
			goto('/search/' + title);
		}
	};
</script>

<div class="flex h-full flex-col items-center justify-center">
	<div class="flex w-full max-w-[40rem] flex-col">
		<h1 class="pb-4 font-mono text-3xl font-semibold">Your way to knowledge</h1>
		<form class="flex w-full gap-2" on:submit|preventDefault={() => search(searchValue)}>
			<Input
				bind:value={searchValue}
				placeholder="Write deck ID or keywords..."
				class="w-full px-4 py-6"
			/>
			<Button class="py-6 py-6" type="submit">
				<Search size={16} strokeWidth={3} />
			</Button>
		</form>
		{#if decks != undefined}
			<div class="grid w-full grid-cols-2 grid-rows-2 gap-2 pt-4">
				{#each decks as deck}
					<div class="aspect-[2/1] w-full">
						<Button
							href={'/deck/' + deck.id}
							variant="outline"
							class="h-full w-full flex-col items-start justify-start gap-2 overflow-hidden rounded-xl p-4"
						>
							<p class="w-full text-wrap">{deck.title}</p>
							<Badge>{deck.cards.length} cards</Badge>
						</Button>
					</div>
				{/each}
			</div>
		{/if}
	</div>
</div>
