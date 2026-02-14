export namespace models {
	
	export class Percentage {
	    reference: string;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new Percentage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.reference = source["reference"];
	        this.size = source["size"];
	    }
	}
	export class BaseInfo {
	    visibilityType: string;
	    xPosition: number;
	    yPosition: number;
	    sizeType: string;
	    absoluteWidth: number;
	    absoluteHeight: number;
	    percentageWidth: Percentage;
	    percentageHeight: Percentage;
	
	    static createFrom(source: any = {}) {
	        return new BaseInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.visibilityType = source["visibilityType"];
	        this.xPosition = source["xPosition"];
	        this.yPosition = source["yPosition"];
	        this.sizeType = source["sizeType"];
	        this.absoluteWidth = source["absoluteWidth"];
	        this.absoluteHeight = source["absoluteHeight"];
	        this.percentageWidth = this.convertValues(source["percentageWidth"], Percentage);
	        this.percentageHeight = this.convertValues(source["percentageHeight"], Percentage);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class PressEvent {
	    autoKeep: boolean;
	    autoClick: boolean;
	    openMenu: boolean;
	    switchTouchMode: boolean;
	    input: boolean;
	    quickInput: boolean;
	    outputText: string;
	    outputKeycodes: number[];
	    bindViewGroup: string[];
	
	    static createFrom(source: any = {}) {
	        return new PressEvent(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.autoKeep = source["autoKeep"];
	        this.autoClick = source["autoClick"];
	        this.openMenu = source["openMenu"];
	        this.switchTouchMode = source["switchTouchMode"];
	        this.input = source["input"];
	        this.quickInput = source["quickInput"];
	        this.outputText = source["outputText"];
	        this.outputKeycodes = source["outputKeycodes"];
	        this.bindViewGroup = source["bindViewGroup"];
	    }
	}
	export class Event {
	    pointerFollow: boolean;
	    Movable: boolean;
	    pressEvent: PressEvent;
	
	    static createFrom(source: any = {}) {
	        return new Event(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pointerFollow = source["pointerFollow"];
	        this.Movable = source["Movable"];
	        this.pressEvent = this.convertValues(source["pressEvent"], PressEvent);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class Button {
	    id: string;
	    text: string;
	    style: string;
	    baseInfo: BaseInfo;
	    event: Event;
	
	    static createFrom(source: any = {}) {
	        return new Button(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.text = source["text"];
	        this.style = source["style"];
	        this.baseInfo = this.convertValues(source["baseInfo"], BaseInfo);
	        this.event = this.convertValues(source["event"], Event);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class ButtonStyle {
	    name: string;
	    textColor: number;
	    textSize: number;
	    strokeColor: number;
	    strokeWidth: number;
	    cornerRadius: number;
	    fillColor: number;
	    textColorPressed: number;
	    textSizePressed: number;
	    strokeColorPressed: number;
	    strokeWidthPressed: number;
	    cornerRadiusPressed: number;
	    fillColorPressed: number;
	
	    static createFrom(source: any = {}) {
	        return new ButtonStyle(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.textColor = source["textColor"];
	        this.textSize = source["textSize"];
	        this.strokeColor = source["strokeColor"];
	        this.strokeWidth = source["strokeWidth"];
	        this.cornerRadius = source["cornerRadius"];
	        this.fillColor = source["fillColor"];
	        this.textColorPressed = source["textColorPressed"];
	        this.textSizePressed = source["textSizePressed"];
	        this.strokeColorPressed = source["strokeColorPressed"];
	        this.strokeWidthPressed = source["strokeWidthPressed"];
	        this.cornerRadiusPressed = source["cornerRadiusPressed"];
	        this.fillColorPressed = source["fillColorPressed"];
	    }
	}
	export class  {
	    locale: string;
	    text: string;
	
	    static createFrom(source: any = {}) {
	        return new (source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.locale = source["locale"];
	        this.text = source["text"];
	    }
	}
	export class Category {
	    id: number;
	    lang: [];
	
	    static createFrom(source: any = {}) {
	        return new Category(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.lang = this.convertValues(source["lang"], );
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class Direction {
	    id: string;
	    style: string;
	    baseInfo: BaseInfo;
	
	    static createFrom(source: any = {}) {
	        return new Direction(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.style = source["style"];
	        this.baseInfo = this.convertValues(source["baseInfo"], BaseInfo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class ViewData {
	    buttonList: Button[];
	    directionList: Direction[];
	
	    static createFrom(source: any = {}) {
	        return new ViewData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.buttonList = this.convertValues(source["buttonList"], Button);
	        this.directionList = this.convertValues(source["directionList"], Direction);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class ViewGroup {
	    id: string;
	    name: string;
	    visibility: string;
	    viewData: ViewData;
	
	    static createFrom(source: any = {}) {
	        return new ViewGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.visibility = source["visibility"];
	        this.viewData = this.convertValues(source["viewData"], ViewData);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class RockerStyle {
	    rockerSize: number;
	    bgCornerRadius: number;
	    bgStrokeWidth: number;
	    bgStrokeColor: number;
	    bgFillColor: number;
	    rockerCornerRadius: number;
	    rockerStrokeWidth: number;
	    rockerStrokeColor: number;
	    rockerFillColor: number;
	
	    static createFrom(source: any = {}) {
	        return new RockerStyle(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.rockerSize = source["rockerSize"];
	        this.bgCornerRadius = source["bgCornerRadius"];
	        this.bgStrokeWidth = source["bgStrokeWidth"];
	        this.bgStrokeColor = source["bgStrokeColor"];
	        this.bgFillColor = source["bgFillColor"];
	        this.rockerCornerRadius = source["rockerCornerRadius"];
	        this.rockerStrokeWidth = source["rockerStrokeWidth"];
	        this.rockerStrokeColor = source["rockerStrokeColor"];
	        this.rockerFillColor = source["rockerFillColor"];
	    }
	}
	export class DirectionStyle {
	    name: string;
	    styleType: string;
	    buttonStyle: ButtonStyle;
	    rockerStyle: RockerStyle;
	
	    static createFrom(source: any = {}) {
	        return new DirectionStyle(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.styleType = source["styleType"];
	        this.buttonStyle = this.convertValues(source["buttonStyle"], ButtonStyle);
	        this.rockerStyle = this.convertValues(source["rockerStyle"], RockerStyle);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	export class ControllerLayout {
	    id: string;
	    name: string;
	    version: string;
	    versionCode: number;
	    author: string;
	    description: string;
	    controllerVersion: number;
	    buttonStyles: ButtonStyle[];
	    directionStyles: DirectionStyle[];
	    viewGroups: ViewGroup[];
	
	    static createFrom(source: any = {}) {
	        return new ControllerLayout(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.version = source["version"];
	        this.versionCode = source["versionCode"];
	        this.author = source["author"];
	        this.description = source["description"];
	        this.controllerVersion = source["controllerVersion"];
	        this.buttonStyles = this.convertValues(source["buttonStyles"], ButtonStyle);
	        this.directionStyles = this.convertValues(source["directionStyles"], DirectionStyle);
	        this.viewGroups = this.convertValues(source["viewGroups"], ViewGroup);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	
	
	
	export class IndexEntry {
	    id: string;
	    lang: string;
	    name: string;
	    introduction: string;
	    device: number[];
	    categories: number[];
	
	    static createFrom(source: any = {}) {
	        return new IndexEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.lang = source["lang"];
	        this.name = source["name"];
	        this.introduction = source["introduction"];
	        this.device = source["device"];
	        this.categories = source["categories"];
	    }
	}
	
	
	export class Version {
	    versionCode: number;
	    versionName: string;
	
	    static createFrom(source: any = {}) {
	        return new Version(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.versionCode = source["versionCode"];
	        this.versionName = source["versionName"];
	    }
	}
	export class RepoVersion {
	    screenshot: number;
	    description: string;
	    author: string;
	    latest: Version;
	    history: Version[];
	
	    static createFrom(source: any = {}) {
	        return new RepoVersion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.screenshot = source["screenshot"];
	        this.description = source["description"];
	        this.author = source["author"];
	        this.latest = this.convertValues(source["latest"], Version);
	        this.history = this.convertValues(source["history"], Version);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

export namespace utils {
	
	export class ParsedPackage {
	    ControllerID: string;
	    VersionCode: number;
	    Layout?: models.ControllerLayout;
	    VersionInfo?: models.RepoVersion;
	    IndexEntry?: models.IndexEntry;
	    IconPath: string;
	    Screenshots: string[];
	    TempDir: string;
	    IsUpdate: boolean;
	    CurrentIndex?: models.IndexEntry;
	
	    static createFrom(source: any = {}) {
	        return new ParsedPackage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ControllerID = source["ControllerID"];
	        this.VersionCode = source["VersionCode"];
	        this.Layout = this.convertValues(source["Layout"], models.ControllerLayout);
	        this.VersionInfo = this.convertValues(source["VersionInfo"], models.RepoVersion);
	        this.IndexEntry = this.convertValues(source["IndexEntry"], models.IndexEntry);
	        this.IconPath = source["IconPath"];
	        this.Screenshots = source["Screenshots"];
	        this.TempDir = source["TempDir"];
	        this.IsUpdate = source["IsUpdate"];
	        this.CurrentIndex = this.convertValues(source["CurrentIndex"], models.IndexEntry);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

