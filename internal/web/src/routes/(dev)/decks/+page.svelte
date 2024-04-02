<script lang="ts">
	import Badge from '$lib/components/ui/badge/badge.svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	let { data } = $props();
</script>

<div class="flex flex-wrap gap-1">
	{#each data.decks as deck}
		<div class="aspect-[2/1] h-32">
			<Button
				href={'/deck/' + deck.id}
				variant="outline"
				class="h-full w-full flex-col items-start justify-start gap-2 overflow-hidden rounded-xl p-4"
			>
				<p class="w-full text-wrap">{deck.title}</p>
				<Badge variant="secondary">{deck.cards.length} cards</Badge>
			</Button>
		</div>
	{/each}
</div>
