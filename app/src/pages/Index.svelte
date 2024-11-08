<script>
    import Layout from "../Layout.svelte";
    import Navbar from "../lib/components/Navbar.svelte";
</script>

<Layout header={Navbar}>
    <main>
        <h1>Index</h1>
        <p>Hello, World!</p>
        <a href="/about">About</a>
    </main>
</Layout>