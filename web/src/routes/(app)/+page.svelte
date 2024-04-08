<script lang="ts">
	import Badge from '$lib/components/ui/badge/badge.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';

	let { data } = $props();
</script>

<div class="flex flex-col items-center justify-center h-full">
	<div class="flex w-full max-w-[40rem] flex-col">
		<h1 class="pb-4 font-mono text-3xl font-semibold">Your way to knowledge</h1>
		<Input placeholder="Write deck ID or keywords..." class="w-full px-4 py-6" />
		<div class="grid w-full grid-cols-2 grid-rows-2 gap-2 pt-4">
			{#each data.decks as deck}
				<div class="aspect-[2/1] w-full">
					<Button
						href={'/deck/' + deck.id}
						variant="outline"
						class="flex-col items-start justify-start w-full h-full gap-2 p-4 overflow-hidden rounded-xl"
					>
						<p class="w-full text-wrap">{deck.title}</p>
						<Badge>{deck.cards.length} cards</Badge>
					</Button>
				</div>
			{/each}
		</div>
	</div>
</div>
