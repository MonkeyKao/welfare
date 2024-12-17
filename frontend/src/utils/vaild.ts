export const isVaildError = (err: unknown, constraints: Record<string, Record<any, any>>): boolean => {
    if (typeof err != "object" || err == null) return false

    return Object.keys(err).every((key) => key in constraints);
}

export const getVaildMessage = (err: unknown): string => {
    if (typeof err == "object" && err != null) {
        for (const value of Object.values(err)) {
            if (Array.isArray(value) && typeof value[0] == "string") {
                return value[0].toLowerCase().replace(Object.keys(err)[0].toLowerCase(), "")
            }
        }
    }
    return "";
}