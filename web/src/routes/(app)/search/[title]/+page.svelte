<script>
	import { CirclePlus } from 'lucide-svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import Badge from '$lib/components/ui/badge/badge.svelte';

	let { data } = $props();
</script>

<div class="flex h-full w-full justify-center">
	<div class="flex w-full max-w-[40rem] flex-col gap-2 pt-6">
		{#if data.decks == undefined}
			<div class="pb-2">
				<h1 class="text-2xl font-medium">No decks found</h1>
				<h2 class="text-muted-foreground font-mono text-sm">
					there is no decks with a simmilar title
				</h2>
			</div>
			<Button href="/create" class="gap-2 p-8" variant="outline">
				<CirclePlus size={16} />
				Create your own
			</Button>
		{:else}
			<div class="pb-2">
				<h1 class="text-2xl font-medium">Related decks</h1>
				<h2 class="text-muted-foreground font-mono text-sm">
					{data.decks.length} decks found
				</h2>
			</div>
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
		{/if}
	</div>
</div>
