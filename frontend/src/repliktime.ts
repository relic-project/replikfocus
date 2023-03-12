import { ref, type Ref } from "vue";

export class ReplikTimeClass {
    private connection: WebSocket | undefined = undefined;
    private reconnect: boolean = true;
    private  maxTries: number = 10;
    public connected: Ref<Boolean> = ref(false);
    public connecting: Ref<Boolean> = ref(false);
    public username: Ref<String> = ref("");
    public users: Ref<String[]> = ref([]);
    public task = ref("");
    public time = ref(Date.now());
    private reconnecting = false;

    connect(): Promise<void> {
        if (this.connected.value && this.connecting.value) {
            return Promise.resolve();
        }
        this.connecting.value = true;

        // get current domain and port
        let domain = window.location.hostname;
        let port = window.location.port;

        // check if env is dev
        if (domain == "localhost" || domain == "") {
            // use localhost and port 8080
            domain = "localhost";
            port = "1323";
        }
        // if https use wss
        let protocol = window.location.protocol == "https:" ? "wss" : "ws";
        this.connection = new WebSocket(`${protocol}://${domain}:${port}/ws`);

        return new Promise((resolve, reject) => {
            this.connection?.addEventListener('open', (ev) => {
                this.connecting.value = false;
                this.connected.value = true;

                this._onopen(ev);
                resolve();
            });
            this.connection?.addEventListener('error', (ev) => {
                this._onerror(ev);
                reject(ev);
            });
            this.connection?.addEventListener('close', (ev) => {
                this.connecting.value = false;
                this.connected.value = false;
                this._onclose(ev);
            });
            this.connection?.addEventListener('message', (ev) => {
                this._onmessage(ev);
            });
        });
    }

    _onopen(ev: Event) {
        this.reconnecting = false;
        console.log("Connected to ReplikTime server", ev);
        this.connection?.send(JSON.stringify({
            type: "register",
            username: this.username.value
        }));
    }

    _onclose(ev: Event) {
        console.log("Disconnected from ReplikTime server", ev);
        this.connection = undefined;
        // exponential backoff for reconnect
        if (this.reconnect && this.reconnecting == false) {
            this.reconnecting = true;
            let delay = 1000;
            let inst = this;
            let tries = 0;
            let backOff = function () {
                if (tries >= inst.maxTries) {
                    console.log("Max reconnect tries reached, giving up");
                    return;
                }
                tries ++;
                setTimeout(() => {
                    console.log("Reconnecting to ReplikTime server (try " + tries + "), delay: " + delay + "ms");
                    inst.connect().catch((err) => {
                        console.log("Error reconnecting to ReplikTime server", err);
                        delay *= 2;
                        backOff();
                    });
                }, delay);
            }
            backOff();
        }
    }

    _onerror(ev: Event) {
        console.log("Error connecting to ReplikTime server", ev);
    }

    _onmessage(ev: MessageEvent) {
        const data = JSON.parse(ev.data);

        if (data.type == "new_timer") {
            console.log("New timer", data);
            this.task.value = data.mode;
            this.time.value = Date.parse(data.expire);
        } else if (data.type == "play_sound") {
            console.log("Play sound", data);
            let audio = new Audio("https://pomofocus.io/audios/alarms/alarm-wood.mp3");
            audio.play();
        } else if (data.type == "register") {
            console.log("Register", data);
            this.username.value = data.username;
        } else if (data.type == "connected_users") {
            console.log("Users", data);
            this.users.value = data.users;
        } else if (data.type == "ping") {

        } else {
            console.log("Message from ReplikTime server", ev);
        }
    }

    disconnect(): Promise<void> {
        this.connection?.close();
        this.connection = undefined;
        this.connecting.value = false;
        this.connected.value = false;
        return Promise.resolve();
    }
}


let ReplikTime = new ReplikTimeClass();
console.log("Connecting to ReplikTime server");
export default ReplikTime;
