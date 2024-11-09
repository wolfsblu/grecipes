<script lang="ts">
    import Layout from "../Layout.svelte";
    import Navbar from "../lib/components/Navbar.svelte";
    import client from "../lib/api/client";

    const fetchRecipes = async () => await client.GET("/recipes")
    let recipeResult = fetchRecipes()
</script>

<Layout header={Navbar}>
    <main>
        <h1>Index</h1>
        <p>Hello, World!</p>
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
    </main>
</Layout>