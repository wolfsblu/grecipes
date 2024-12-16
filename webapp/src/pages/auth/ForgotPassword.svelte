<script lang="ts">
    import Button from "../../lib/components/forms/Button.svelte";
    import Input from "../../lib/components/forms/Input.svelte";
    import Layout from "../../Layout.svelte";
    import LoginIcon from "../../lib/icons/Login.svelte";
    import Navbar from "../../lib/components/navigation/Navbar.svelte";
    import t from "../../lib/i18n/i18n.svelte.js"
    import {createRouter} from "../../lib/router.svelte.js";
    import {createUser} from "../../lib/auth/user.svelte.js";
    import food from "../../assets/images/login.jpg"

    const router = createRouter()
    const user = createUser()

    let credentials: Credentials = $state({
        email: "",
        password: "",
    })

    const handleSubmit = async (e: SubmitEvent) => {
        e.preventDefault()
        try {
            await user.login(credentials)
            router.redirectToNext()
        } catch (e) {
            // TODO: Show error toast
        }
    }
</script>

<Layout Header={Navbar}>
    <div class="flex flex-col flex-wrap md:h-full items-center justify-center">
        <div class="md:grid grid-cols-[1fr_2fr] md:shadow-lg md:w-1/2">
            <div class="bg-center bg-cover rounded-l-md"
                 style="background-image: url({food})">
            </div>
            <form class="md:border border-gray-200 grid grid-rows-[auto_min-content] rounded-r-md" onsubmit="{handleSubmit}">
                <div class="p-6">
                    <h1 class="font-light mb-3 text-3xl">{t("forgot-password.title")}</h1>
                    <p class="mb-3">{t("forgot-password.subtitle")}</p>
                    <Input label={t("login.labels.email")} type="email" bind:value={credentials.email} required={true}/>
                    <p class="text-sm">
                        {@html t("forgot-password.help.email", {
                            class: "hover:no-underline hover:text-blue-500 text-blue-600 underline",
                            url: "/"
                        })}
                    </p>
                </div>

                <div class="p-6 pt-0">
                    <Button type="submit" class="mt-3" icon={LoginIcon}>
                        {t("forgot-password.actions.submit")}
                    </Button>
                </div>
            </form>
        </div>
    </div>
</Layout>
