let profile = $state(null)

export const createUser = () => {
    const login = (user: any) => {
        profile = user
    }

    return {
        get profile() {
            return profile
        },
        login
    }
}