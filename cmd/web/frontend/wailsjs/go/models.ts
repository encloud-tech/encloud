export namespace types {
	
	
	
	export class FetchKeysResponse {
	    publicKey: string;
	    files: number;
	
	    static createFrom(source: any = {}) {
	        return new FetchKeysResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.publicKey = source["publicKey"];
	        this.files = source["files"];
	    }
	}
	export class FileMetadata {
	    uuid: string;
	    publicKey: string;
	    md5Hash: string;
	    timestamp: number;
	    uploadedAt: string;
	    name: string;
	    size: number;
	    fileType: string;
	    cid: string[];
	    dek: number[];
	    dekType: string;
	    kekType: string;
	
	    static createFrom(source: any = {}) {
	        return new FileMetadata(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uuid = source["uuid"];
	        this.publicKey = source["publicKey"];
	        this.md5Hash = source["md5Hash"];
	        this.timestamp = source["timestamp"];
	        this.uploadedAt = source["uploadedAt"];
	        this.name = source["name"];
	        this.size = source["size"];
	        this.fileType = source["fileType"];
	        this.cid = source["cid"];
	        this.dek = source["dek"];
	        this.dekType = source["dekType"];
	        this.kekType = source["kekType"];
	    }
	}
	
	export class Keys {
	    PublicKey: string;
	    PrivateKey: string;
	
	    static createFrom(source: any = {}) {
	        return new Keys(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.PublicKey = source["PublicKey"];
	        this.PrivateKey = source["PrivateKey"];
	    }
	}
	
	
	
	

}

