<script lang="ts">
    import Layout from "../Layout.svelte";
    import Navbar from "../lib/components/navigation/Navbar.svelte";
    import {fetchRecipes} from "../lib/api/client";
    import t, {localeSwitcher} from "../lib/i18n/i18n.svelte";

    let recipeResult = fetchRecipes()
</script>

<Layout Header={Navbar}>
    <h1>Index</h1>
    <select bind:value={localeSwitcher.locale}>
        <option value="en">EN</option>
        <option value="de">DE</option>
    </select>
    <a href="/about">About</a>
    {#await recipeResult}
        <p>
            Loading recipes...
        </p>
    {:then recipes}
        {#if !recipes.data || recipes.data.length <= 0}
            <p>No recipes found</p>
        {:else}
            <ul>
                {#each recipes.data as recipe}
                    <li>{recipe.name}</li>
                {/each}
            </ul>
        {/if}
    {:catch error}
        <p>Failed to load recipes</p>
    {/await}
</Layout>