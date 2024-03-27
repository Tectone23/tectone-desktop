export namespace main {
	
	export class Project {
	    key_hash: string;
	    id: string;
	    name: string;
	    path: string;
	    sdk_path: string;
	    sandbox_path: string;
	    containers: model.Container[];
	    testnet: model.TestNet;
	    devnet: model.DevNet;
	    mainnet: model.MainNet;
	    created_at: number;
	
	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key_hash = source["key_hash"];
	        this.id = source["id"];
	        this.name = source["name"];
	        this.path = source["path"];
	        this.sdk_path = source["sdk_path"];
	        this.sandbox_path = source["sandbox_path"];
	        this.containers = this.convertValues(source["containers"], model.Container);
	        this.testnet = this.convertValues(source["testnet"], model.TestNet);
	        this.devnet = this.convertValues(source["devnet"], model.DevNet);
	        this.mainnet = this.convertValues(source["mainnet"], model.MainNet);
	        this.created_at = source["created_at"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace model {
	
	export class State {
	    algo: number;
	    onl?: number;
	    sel?: string;
	    vote?: string;
	    voteKD?: number;
	    voteLst?: number;
	
	    static createFrom(source: any = {}) {
	        return new State(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.algo = source["algo"];
	        this.onl = source["onl"];
	        this.sel = source["sel"];
	        this.vote = source["vote"];
	        this.voteKD = source["voteKD"];
	        this.voteLst = source["voteLst"];
	    }
	}
	export class Allocation {
	    addr: string;
	    comment: string;
	    state: State;
	
	    static createFrom(source: any = {}) {
	        return new Allocation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.addr = source["addr"];
	        this.comment = source["comment"];
	        this.state = this.convertValues(source["state"], State);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class AppInfo {
	    version: string;
	    commit_hash: string;
	    build_date: string;
	
	    static createFrom(source: any = {}) {
	        return new AppInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.commit_hash = source["commit_hash"];
	        this.build_date = source["build_date"];
	    }
	}
	export class Container {
	    id: string;
	    name: string;
	    size: number;
	    status: string;
	
	    static createFrom(source: any = {}) {
	        return new Container(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.size = source["size"];
	        this.status = source["status"];
	    }
	}
	export class Genesis {
	    alloc: Allocation[];
	    fees: string;
	    id: string;
	    network: string;
	    proto?: string;
	    rwd: string;
	    timestamp?: number;
	
	    static createFrom(source: any = {}) {
	        return new Genesis(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.alloc = this.convertValues(source["alloc"], Allocation);
	        this.fees = source["fees"];
	        this.id = source["id"];
	        this.network = source["network"];
	        this.proto = source["proto"];
	        this.rwd = source["rwd"];
	        this.timestamp = source["timestamp"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class DevNet {
	    genesis: Genesis;
	
	    static createFrom(source: any = {}) {
	        return new DevNet(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.genesis = this.convertValues(source["genesis"], Genesis);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Feedback {
	    name: string;
	    email: string;
	    filepath: string;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new Feedback(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.email = source["email"];
	        this.filepath = source["filepath"];
	        this.message = source["message"];
	    }
	}
	
	export class MainNet {
	    genesis: Genesis;
	
	    static createFrom(source: any = {}) {
	        return new MainNet(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.genesis = this.convertValues(source["genesis"], Genesis);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TestNet {
	    genesis: Genesis;
	
	    static createFrom(source: any = {}) {
	        return new TestNet(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.genesis = this.convertValues(source["genesis"], Genesis);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Project {
	    key_hash: string;
	    id: string;
	    name: string;
	    path: string;
	    sdk_path: string;
	    sandbox_path: string;
	    containers: Container[];
	    testnet: TestNet;
	    devnet: DevNet;
	    mainnet: MainNet;
	    created_at: number;
	
	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key_hash = source["key_hash"];
	        this.id = source["id"];
	        this.name = source["name"];
	        this.path = source["path"];
	        this.sdk_path = source["sdk_path"];
	        this.sandbox_path = source["sandbox_path"];
	        this.containers = this.convertValues(source["containers"], Container);
	        this.testnet = this.convertValues(source["testnet"], TestNet);
	        this.devnet = this.convertValues(source["devnet"], DevNet);
	        this.mainnet = this.convertValues(source["mainnet"], MainNet);
	        this.created_at = source["created_at"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

