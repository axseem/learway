<script lang="ts">
	import Badge from '$lib/components/ui/badge/badge.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';

	let { data } = $props();
	$effect(() => {});
</script>

<div class="flex h-full flex-col items-center justify-center">
	<div class="flex w-full max-w-[40rem] flex-col">
		<h1 class="pb-2 font-mono text-3xl font-semibold">Personal way to learn</h1>
		<Input placeholder="Write deck ID or keywords..." class="w-full px-4 py-6" />
		<div class="grid w-full grid-cols-2 grid-rows-2 gap-2 pt-4">
			{#each data.decks as deck}
				<div class="aspect-[2/1] w-full">
					<Button
						href={'deck/' + deck.id}
						variant="outline"
						class="h-full w-full flex-col items-start justify-start gap-2 overflow-hidden rounded-xl p-4"
					>
						<p class="w-full text-wrap">{deck.title}</p>
						<Badge variant="secondary">{deck.cards.length} cards</Badge>
					</Button>
				</div>
			{/each}
		</div>
	</div>
</div>
