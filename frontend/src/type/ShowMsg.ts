export enum AlertColor { "error", "success", "info","waring" }

export type showMsgFunction = (msg: string, color?: AlertColor, clickFunction?: () => void) => void;