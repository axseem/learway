<script>
	let { children } = $props();
</script>

<div class="flex items-center justify-center w-screen h-screen p-4">
	<div class="flex flex-col w-full max-w-80">
		{@render children()}
	</div>
</div>
