export namespace models {
	
	export class User {
	    name: string;
	    role: string;
	    email: string;
	    phone: string;
	    address: string;
	    city: string;
	    state: string;
	    zip: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.role = source["role"];
	        this.email = source["email"];
	        this.phone = source["phone"];
	        this.address = source["address"];
	        this.city = source["city"];
	        this.state = source["state"];
	        this.zip = source["zip"];
	    }
	}

}

