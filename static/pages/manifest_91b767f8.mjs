import '@astrojs/internal-helpers/path';
import 'cookie';
import 'kleur/colors';
import 'string-width';
import 'mime';
import 'html-escaper';
import 'clsx';
import './chunks/astro_96bdb4d7.mjs';
import { compile } from 'path-to-regexp';

if (typeof process !== "undefined") {
  let proc = process;
  if ("argv" in proc && Array.isArray(proc.argv)) {
    if (proc.argv.includes("--verbose")) ; else if (proc.argv.includes("--silent")) ; else ;
  }
}

new TextEncoder();

function getRouteGenerator(segments, addTrailingSlash) {
  const template = segments.map((segment) => {
    return "/" + segment.map((part) => {
      if (part.spread) {
        return `:${part.content.slice(3)}(.*)?`;
      } else if (part.dynamic) {
        return `:${part.content}`;
      } else {
        return part.content.normalize().replace(/\?/g, "%3F").replace(/#/g, "%23").replace(/%5B/g, "[").replace(/%5D/g, "]").replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
      }
    }).join("");
  }).join("");
  let trailing = "";
  if (addTrailingSlash === "always" && segments.length) {
    trailing = "/";
  }
  const toPath = compile(template + trailing);
  return toPath;
}

function deserializeRouteData(rawRouteData) {
  return {
    route: rawRouteData.route,
    type: rawRouteData.type,
    pattern: new RegExp(rawRouteData.pattern),
    params: rawRouteData.params,
    component: rawRouteData.component,
    generate: getRouteGenerator(rawRouteData.segments, rawRouteData._meta.trailingSlash),
    pathname: rawRouteData.pathname || void 0,
    segments: rawRouteData.segments,
    prerender: rawRouteData.prerender,
    redirect: rawRouteData.redirect,
    redirectRoute: rawRouteData.redirectRoute ? deserializeRouteData(rawRouteData.redirectRoute) : void 0,
    fallbackRoutes: rawRouteData.fallbackRoutes.map((fallback) => {
      return deserializeRouteData(fallback);
    })
  };
}

function deserializeManifest(serializedManifest) {
  const routes = [];
  for (const serializedRoute of serializedManifest.routes) {
    routes.push({
      ...serializedRoute,
      routeData: deserializeRouteData(serializedRoute.routeData)
    });
    const route = serializedRoute;
    route.routeData = deserializeRouteData(serializedRoute.routeData);
  }
  const assets = new Set(serializedManifest.assets);
  const componentMetadata = new Map(serializedManifest.componentMetadata);
  const clientDirectives = new Map(serializedManifest.clientDirectives);
  return {
    ...serializedManifest,
    assets,
    componentMetadata,
    clientDirectives,
    routes
  };
}

const manifest = deserializeManifest({"adapterName":"","routes":[{"file":"","links":[],"scripts":[],"styles":[{"type":"inline","content":"*{font-family:Lato;margin:0;padding:0;box-sizing:border-box}html{background-color:#222;color:#ddf;display:grid;place-items:center;width:100vw;height:100vh}main{padding:32px;margin:16px;max-width:500px;max-height:500px;background:#333333;border-radius:8px}a{color:#99f}h1{margin-bottom:8px;text-align:center}p{margin:4px 0}img{border-radius:4px;margin:4px 0}\n"}],"routeData":{"route":"/","type":"page","pattern":"^\\/$","segments":[],"params":[],"component":"src/pages/index.astro","pathname":"/","prerender":false,"fallbackRoutes":[],"_meta":{"trailingSlash":"ignore"}}},{"file":"","links":[],"scripts":[],"styles":[{"type":"inline","content":"*{font-family:Lato;margin:0;padding:0;box-sizing:border-box}html{background-color:#222;color:#ddf;display:grid;place-items:center;width:100vw;height:100vh}main{padding:32px;margin:16px;max-width:500px;max-height:500px;background:#333333;border-radius:8px}a{color:#99f}h1{margin-bottom:8px;text-align:center}p{margin:4px 0}img{border-radius:4px;margin:4px 0}\n"}],"routeData":{"route":"/checkmate","type":"page","pattern":"^\\/checkmate\\/?$","segments":[[{"content":"checkmate","dynamic":false,"spread":false}]],"params":[],"component":"src/pages/checkmate.astro","pathname":"/checkmate","prerender":false,"fallbackRoutes":[],"_meta":{"trailingSlash":"ignore"}}},{"file":"","links":[],"scripts":[],"styles":[{"type":"inline","content":"*{font-family:Lato;margin:0;padding:0;box-sizing:border-box}html{background-color:#222;color:#ddf;display:grid;place-items:center;width:100vw;height:100vh}main{padding:32px;margin:16px;max-width:500px;max-height:500px;background:#333333;border-radius:8px}a{color:#99f}h1{margin-bottom:8px;text-align:center}p{margin:4px 0}img{border-radius:4px;margin:4px 0}\ndiv{overflow:hidden}img[alt]{margin-left:-16px;text-indent:16px}\n"}],"routeData":{"route":"/quark","type":"page","pattern":"^\\/quark\\/?$","segments":[[{"content":"quark","dynamic":false,"spread":false}]],"params":[],"component":"src/pages/quark.astro","pathname":"/quark","prerender":false,"fallbackRoutes":[],"_meta":{"trailingSlash":"ignore"}}},{"file":"","links":[],"scripts":[{"type":"inline","value":"const b=window,n=b[atob(\"ZGV2aWNlUGl4ZWxSYXRpbw==\")];function o(){const l=b[atob(\"ZGV2aWNlUGl4ZWxSYXRpbw==\")]>=n*1.5,c=document[atob(\"cXVlcnlTZWxlY3Rvcg==\")](\"p\");l?c[atob(\"dGV4dENvbnRlbnQ=\")]=atob(\"cXVhcms=\"):c[atob(\"dGV4dENvbnRlbnQ=\")]=atob(\"dGhlIGFuc3dlciBpcyB0aW55LiB5b3Ugd2lsbCBoYXZlIHRvIHVzZSBhIG1hZ25pZnlpbmcgZ2xhc3Mu\")}b.addEventListener(atob(\"cmVzaXpl\"),o);\n"}],"styles":[{"type":"inline","content":"*{font-family:Lato;margin:0;padding:0;box-sizing:border-box}html{background-color:#222;color:#ddf;display:grid;place-items:center;width:100vw;height:100vh}main{padding:32px;margin:16px;max-width:500px;max-height:500px;background:#333333;border-radius:8px}a{color:#99f}h1{margin-bottom:8px;text-align:center}p{margin:4px 0}img{border-radius:4px;margin:4px 0}\n"}],"routeData":{"route":"/here","type":"page","pattern":"^\\/here\\/?$","segments":[[{"content":"here","dynamic":false,"spread":false}]],"params":[],"component":"src/pages/here.astro","pathname":"/here","prerender":false,"fallbackRoutes":[],"_meta":{"trailingSlash":"ignore"}}},{"file":"","links":[],"scripts":[],"styles":[{"type":"inline","content":"*{font-family:Lato;margin:0;padding:0;box-sizing:border-box}html{background-color:#222;color:#ddf;display:grid;place-items:center;width:100vw;height:100vh}main{padding:32px;margin:16px;max-width:500px;max-height:500px;background:#333333;border-radius:8px}a{color:#99f}h1{margin-bottom:8px;text-align:center}p{margin:4px 0}img{border-radius:4px;margin:4px 0}\n"}],"routeData":{"route":"/move","type":"page","pattern":"^\\/move\\/?$","segments":[[{"content":"move","dynamic":false,"spread":false}]],"params":[],"component":"src/pages/move.astro","pathname":"/move","prerender":false,"fallbackRoutes":[],"_meta":{"trailingSlash":"ignore"}}},{"file":"","links":[],"scripts":[],"styles":[{"type":"inline","content":"*{font-family:Lato;margin:0;padding:0;box-sizing:border-box}html{background-color:#222;color:#ddf;display:grid;place-items:center;width:100vw;height:100vh}main{padding:32px;margin:16px;max-width:500px;max-height:500px;background:#333333;border-radius:8px}a{color:#99f}h1{margin-bottom:8px;text-align:center}p{margin:4px 0}img{border-radius:4px;margin:4px 0}\n"}],"routeData":{"route":"/qxf7","type":"page","pattern":"^\\/qxf7\\/?$","segments":[[{"content":"qxf7","dynamic":false,"spread":false}]],"params":[],"component":"src/pages/qxf7.astro","pathname":"/qxf7","prerender":false,"fallbackRoutes":[],"_meta":{"trailingSlash":"ignore"}}},{"file":"","links":[],"scripts":[],"styles":[{"type":"inline","content":"*{font-family:Lato;margin:0;padding:0;box-sizing:border-box}html{background-color:#222;color:#ddf;display:grid;place-items:center;width:100vw;height:100vh}main{padding:32px;margin:16px;max-width:500px;max-height:500px;background:#333333;border-radius:8px}a{color:#99f}h1{margin-bottom:8px;text-align:center}p{margin:4px 0}img{border-radius:4px;margin:4px 0}\n"}],"routeData":{"route":"/404","type":"page","pattern":"^\\/404\\/?$","segments":[[{"content":"404","dynamic":false,"spread":false}]],"params":[],"component":"src/pages/404.astro","pathname":"/404","prerender":false,"fallbackRoutes":[],"_meta":{"trailingSlash":"ignore"}}},{"file":"","links":[],"scripts":[],"styles":[{"type":"inline","content":"*{font-family:Lato;margin:0;padding:0;box-sizing:border-box}html{background-color:#222;color:#ddf;display:grid;place-items:center;width:100vw;height:100vh}main{padding:32px;margin:16px;max-width:500px;max-height:500px;background:#333333;border-radius:8px}a{color:#99f}h1{margin-bottom:8px;text-align:center}p{margin:4px 0}img{border-radius:4px;margin:4px 0}\nspan[data-astro-cid-i4lgtgsn]{color:#bbf}\n"}],"routeData":{"route":"/one","type":"page","pattern":"^\\/one\\/?$","segments":[[{"content":"one","dynamic":false,"spread":false}]],"params":[],"component":"src/pages/one.astro","pathname":"/one","prerender":false,"fallbackRoutes":[],"_meta":{"trailingSlash":"ignore"}}},{"file":"","links":[],"scripts":[],"styles":[{"type":"inline","content":"*{font-family:Lato;margin:0;padding:0;box-sizing:border-box}html{background-color:#222;color:#ddf;display:grid;place-items:center;width:100vw;height:100vh}main{padding:32px;margin:16px;max-width:500px;max-height:500px;background:#333333;border-radius:8px}a{color:#99f}h1{margin-bottom:8px;text-align:center}p{margin:4px 0}img{border-radius:4px;margin:4px 0}\n"}],"routeData":{"route":"/two","type":"page","pattern":"^\\/two\\/?$","segments":[[{"content":"two","dynamic":false,"spread":false}]],"params":[],"component":"src/pages/two.astro","pathname":"/two","prerender":false,"fallbackRoutes":[],"_meta":{"trailingSlash":"ignore"}}},{"file":"","links":[],"scripts":[],"styles":[{"type":"inline","content":"*{font-family:Lato;margin:0;padding:0;box-sizing:border-box}html{background-color:#222;color:#ddf;display:grid;place-items:center;width:100vw;height:100vh}main{padding:32px;margin:16px;max-width:500px;max-height:500px;background:#333333;border-radius:8px}a{color:#99f}h1{margin-bottom:8px;text-align:center}p{margin:4px 0}img{border-radius:4px;margin:4px 0}\n"}],"routeData":{"route":"/b/checkmate","type":"page","pattern":"^\\/b\\/checkmate\\/?$","segments":[[{"content":"b","dynamic":false,"spread":false}],[{"content":"checkmate","dynamic":false,"spread":false}]],"params":[],"component":"src/pages/b/checkmate.astro","pathname":"/b/checkmate","prerender":false,"fallbackRoutes":[],"_meta":{"trailingSlash":"ignore"}}}],"base":"/","trailingSlash":"ignore","compressHTML":true,"componentMetadata":[["/home/bieniucieniu/dev/noestabien/web/astro/src/pages/404.astro",{"propagation":"none","containsHead":true}],["/home/bieniucieniu/dev/noestabien/web/astro/src/pages/b/checkmate.astro",{"propagation":"none","containsHead":true}],["/home/bieniucieniu/dev/noestabien/web/astro/src/pages/checkmate.astro",{"propagation":"none","containsHead":true}],["/home/bieniucieniu/dev/noestabien/web/astro/src/pages/here.astro",{"propagation":"none","containsHead":true}],["/home/bieniucieniu/dev/noestabien/web/astro/src/pages/index.astro",{"propagation":"none","containsHead":true}],["/home/bieniucieniu/dev/noestabien/web/astro/src/pages/move.astro",{"propagation":"none","containsHead":true}],["/home/bieniucieniu/dev/noestabien/web/astro/src/pages/one.astro",{"propagation":"none","containsHead":true}],["/home/bieniucieniu/dev/noestabien/web/astro/src/pages/quark.astro",{"propagation":"none","containsHead":true}],["/home/bieniucieniu/dev/noestabien/web/astro/src/pages/qxf7.astro",{"propagation":"none","containsHead":true}],["/home/bieniucieniu/dev/noestabien/web/astro/src/pages/two.astro",{"propagation":"none","containsHead":true}]],"renderers":[],"clientDirectives":[["idle","(()=>{var i=t=>{let e=async()=>{await(await t())()};\"requestIdleCallback\"in window?window.requestIdleCallback(e):setTimeout(e,200)};(self.Astro||(self.Astro={})).idle=i;window.dispatchEvent(new Event(\"astro:idle\"));})();"],["load","(()=>{var e=async t=>{await(await t())()};(self.Astro||(self.Astro={})).load=e;window.dispatchEvent(new Event(\"astro:load\"));})();"],["media","(()=>{var s=(i,t)=>{let a=async()=>{await(await i())()};if(t.value){let e=matchMedia(t.value);e.matches?a():e.addEventListener(\"change\",a,{once:!0})}};(self.Astro||(self.Astro={})).media=s;window.dispatchEvent(new Event(\"astro:media\"));})();"],["only","(()=>{var e=async t=>{await(await t())()};(self.Astro||(self.Astro={})).only=e;window.dispatchEvent(new Event(\"astro:only\"));})();"],["visible","(()=>{var r=(i,c,s)=>{let n=async()=>{await(await i())()},t=new IntersectionObserver(e=>{for(let o of e)if(o.isIntersecting){t.disconnect(),n();break}});for(let e of s.children)t.observe(e)};(self.Astro||(self.Astro={})).visible=r;window.dispatchEvent(new Event(\"astro:visible\"));})();"]],"entryModules":{"\u0000@astro-page:src/pages/index@_@astro":"pages/index.astro.mjs","\u0000@astro-page:src/pages/checkmate@_@astro":"pages/checkmate.astro.mjs","\u0000@astro-page:src/pages/quark@_@astro":"pages/quark.astro.mjs","\u0000@astro-page:src/pages/here@_@astro":"pages/here.astro.mjs","\u0000@astro-page:src/pages/move@_@astro":"pages/move.astro.mjs","\u0000@astro-page:src/pages/qxf7@_@astro":"pages/qxf7.astro.mjs","\u0000@astro-page:src/pages/404@_@astro":"pages/404.astro.mjs","\u0000@astro-page:src/pages/one@_@astro":"pages/one.astro.mjs","\u0000@astro-page:src/pages/two@_@astro":"pages/two.astro.mjs","\u0000@astro-page:src/pages/b/checkmate@_@astro":"pages/b/checkmate.astro.mjs","\u0000@astro-renderers":"renderers.mjs","\u0000empty-middleware":"_empty-middleware.mjs","/src/pages/here.astro":"chunks/pages/here_58e06b1e.mjs","/src/pages/index.astro":"chunks/pages/index_6db04dc3.mjs","/src/pages/move.astro":"chunks/pages/move_d06adea2.mjs","/src/pages/one.astro":"chunks/pages/one_eb769f31.mjs","/src/pages/quark.astro":"chunks/pages/quark_e110a010.mjs","/src/pages/qxf7.astro":"chunks/pages/qxf7_a9c32cf9.mjs","/src/pages/two.astro":"chunks/pages/two_f0ff93da.mjs","\u0000@astrojs-manifest":"manifest_91b767f8.mjs","/astro/hoisted.js?q=0":"_astro/hoisted.10312f6f.js","astro:scripts/before-hydration.js":""},"assets":[]});

export { manifest };
