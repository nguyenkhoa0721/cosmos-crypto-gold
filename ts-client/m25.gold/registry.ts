import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgTransferTo } from "./types/gold/tx";
import { MsgTransfer } from "./types/gold/tx";
import { MsgMint } from "./types/gold/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/m25.gold.MsgTransferTo", MsgTransferTo],
    ["/m25.gold.MsgTransfer", MsgTransfer],
    ["/m25.gold.MsgMint", MsgMint],
    
];

export { msgTypes }