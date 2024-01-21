import { addedLog, constructedLog, removedLog } from "../helpers/helper";

const ELEMENT = "lazy-comp-two";

export default class LazyCompTwo extends HTMLElement {
  constructor() {
    super();
    console.log(constructedLog(ELEMENT));
  }

  connectedCallback() {
    console.log(addedLog(ELEMENT));
  }

  disconnectedCallback() {
    console.log(removedLog(ELEMENT));
  }
}

customElements.define(ELEMENT, LazyCompTwo);
