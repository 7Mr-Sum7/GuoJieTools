export namespace main {
	
	export class RespDate {
	    code: number;
	    data: any;
	    msg: string;
	
	    static createFrom(source: any = {}) {
	        return new RespDate(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.data = source["data"];
	        this.msg = source["msg"];
	    }
	}

}

