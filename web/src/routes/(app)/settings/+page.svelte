<script lang="ts">
	let darkMode = $state(true);

	function switchTheme() {
		darkMode = !darkMode;

		localStorage.setItem('theme', darkMode ? 'dark' : 'light');

		darkMode
			? document.documentElement.classList.add('dark')
			: document.documentElement.classList.remove('dark');
	}

	if (
		localStorage.theme === 'dark' ||
		(!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)
	) {
		document.documentElement.classList.add('dark');
		darkMode = true;
	} else {
		document.documentElement.classList.remove('dark');
		darkMode = false;
	}
</script>

<div class="flex h-full w-full justify-center">
	<div class="flex w-full max-w-[40rem] flex-col gap-4 pt-6">
		<div class="flex items-center gap-2">
			<input
				on:change={switchTheme}
				bind:checked={darkMode}
				type="checkbox"
				name="dark"
				id="dark"
			/>
			<label for="dark">Dark mode</label>
		</div>
	</div>
</div>
