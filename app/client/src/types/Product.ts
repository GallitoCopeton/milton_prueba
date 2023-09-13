import {Model} from "@vuex-orm/core";

export default class Product extends Model {
    static entity = "products";

    id!: number;
    name!: string;
    price!: number;
    rank!: number;

    static fields() {
        return {
            id: this.number(0),
            name: this.string(""),
            price: this.number(0),
            rank: this.number(0),
        };
    }
}