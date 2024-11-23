<script>
    import Plus from "../../icons/Plus.svelte";
    import logo from "../../../assets/images/logo.svg"
    import NavButton from "./NavButton.svelte";
</script>

<div class="flex items-center justify-between bg-orange-200 px-3 py-2">
    <a class="flex items-center gap-x-3" href="/">
        <span class="block bg-orange-50 rounded-full w-11 h-11 p-2">
            <img src={logo} alt="Logo">
        </span>
        <h1 class="text-2xl text-alice font-light text-orange-800">Recipes</h1>
    </a>

    <NavButton text="New Recipe" href="/recipes/create" icon={Plus} />
</div>