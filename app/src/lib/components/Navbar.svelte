<script>
    import Plus from "../icons/Plus.svelte";
    import logo from "../../assets/images/logo.svg"
</script>

<div class="flex items-center justify-between bg-gray-300 px-3 py-2">
    <a class="flex items-center gap-x-3" href="/">
        <span class="block bg-alice rounded-full w-11 h-11 p-2">
            <img src={logo} alt="Logo">
        </span>
        <h1 class="text-2xl text-alice font-light">Recipes</h1>
    </a>

    <a class="bg-dogwood flex items-center rounded-full sm:rounded-lg p-3 gap-x-2" href="/recipes/create/">
        <span class="h-4 w-4 sm:hidden">
            <Plus/>
        </span>
        <span class="hidden sm:inline">New Recipe</span>
    </a>

</div>