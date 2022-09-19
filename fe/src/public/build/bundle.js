
(function(l, r) { if (!l || l.getElementById('livereloadscript')) return; r = l.createElement('script'); r.async = 1; r.src = '//' + (self.location.host || 'localhost').split(':')[0] + ':35729/livereload.js?snipver=1'; r.id = 'livereloadscript'; l.getElementsByTagName('head')[0].appendChild(r) })(self.document);
var app = (function () {
    'use strict';

    function noop() { }
    const identity = x => x;
    function assign(tar, src) {
        // @ts-ignore
        for (const k in src)
            tar[k] = src[k];
        return tar;
    }
    function add_location(element, file, line, column, char) {
        element.__svelte_meta = {
            loc: { file, line, column, char }
        };
    }
    function run(fn) {
        return fn();
    }
    function blank_object() {
        return Object.create(null);
    }
    function run_all(fns) {
        fns.forEach(run);
    }
    function is_function(thing) {
        return typeof thing === 'function';
    }
    function safe_not_equal(a, b) {
        return a != a ? b == b : a !== b || ((a && typeof a === 'object') || typeof a === 'function');
    }
    function is_empty(obj) {
        return Object.keys(obj).length === 0;
    }
    function create_slot(definition, ctx, $$scope, fn) {
        if (definition) {
            const slot_ctx = get_slot_context(definition, ctx, $$scope, fn);
            return definition[0](slot_ctx);
        }
    }
    function get_slot_context(definition, ctx, $$scope, fn) {
        return definition[1] && fn
            ? assign($$scope.ctx.slice(), definition[1](fn(ctx)))
            : $$scope.ctx;
    }
    function get_slot_changes(definition, $$scope, dirty, fn) {
        if (definition[2] && fn) {
            const lets = definition[2](fn(dirty));
            if ($$scope.dirty === undefined) {
                return lets;
            }
            if (typeof lets === 'object') {
                const merged = [];
                const len = Math.max($$scope.dirty.length, lets.length);
                for (let i = 0; i < len; i += 1) {
                    merged[i] = $$scope.dirty[i] | lets[i];
                }
                return merged;
            }
            return $$scope.dirty | lets;
        }
        return $$scope.dirty;
    }
    function update_slot_base(slot, slot_definition, ctx, $$scope, slot_changes, get_slot_context_fn) {
        if (slot_changes) {
            const slot_context = get_slot_context(slot_definition, ctx, $$scope, get_slot_context_fn);
            slot.p(slot_context, slot_changes);
        }
    }
    function get_all_dirty_from_scope($$scope) {
        if ($$scope.ctx.length > 32) {
            const dirty = [];
            const length = $$scope.ctx.length / 32;
            for (let i = 0; i < length; i++) {
                dirty[i] = -1;
            }
            return dirty;
        }
        return -1;
    }

    const is_client = typeof window !== 'undefined';
    let now = is_client
        ? () => window.performance.now()
        : () => Date.now();
    let raf = is_client ? cb => requestAnimationFrame(cb) : noop;

    const tasks = new Set();
    function run_tasks(now) {
        tasks.forEach(task => {
            if (!task.c(now)) {
                tasks.delete(task);
                task.f();
            }
        });
        if (tasks.size !== 0)
            raf(run_tasks);
    }
    /**
     * Creates a new task that runs on each raf frame
     * until it returns a falsy value or is aborted
     */
    function loop(callback) {
        let task;
        if (tasks.size === 0)
            raf(run_tasks);
        return {
            promise: new Promise(fulfill => {
                tasks.add(task = { c: callback, f: fulfill });
            }),
            abort() {
                tasks.delete(task);
            }
        };
    }
    function append(target, node) {
        target.appendChild(node);
    }
    function get_root_for_style(node) {
        if (!node)
            return document;
        const root = node.getRootNode ? node.getRootNode() : node.ownerDocument;
        if (root && root.host) {
            return root;
        }
        return node.ownerDocument;
    }
    function append_empty_stylesheet(node) {
        const style_element = element('style');
        append_stylesheet(get_root_for_style(node), style_element);
        return style_element.sheet;
    }
    function append_stylesheet(node, style) {
        append(node.head || node, style);
        return style.sheet;
    }
    function insert(target, node, anchor) {
        target.insertBefore(node, anchor || null);
    }
    function detach(node) {
        node.parentNode.removeChild(node);
    }
    function destroy_each(iterations, detaching) {
        for (let i = 0; i < iterations.length; i += 1) {
            if (iterations[i])
                iterations[i].d(detaching);
        }
    }
    function element(name) {
        return document.createElement(name);
    }
    function svg_element(name) {
        return document.createElementNS('http://www.w3.org/2000/svg', name);
    }
    function text(data) {
        return document.createTextNode(data);
    }
    function space() {
        return text(' ');
    }
    function listen(node, event, handler, options) {
        node.addEventListener(event, handler, options);
        return () => node.removeEventListener(event, handler, options);
    }
    function prevent_default(fn) {
        return function (event) {
            event.preventDefault();
            // @ts-ignore
            return fn.call(this, event);
        };
    }
    function attr(node, attribute, value) {
        if (value == null)
            node.removeAttribute(attribute);
        else if (node.getAttribute(attribute) !== value)
            node.setAttribute(attribute, value);
    }
    function children(element) {
        return Array.from(element.childNodes);
    }
    function set_input_value(input, value) {
        input.value = value == null ? '' : value;
    }
    function set_style(node, key, value, important) {
        if (value === null) {
            node.style.removeProperty(key);
        }
        else {
            node.style.setProperty(key, value, important ? 'important' : '');
        }
    }
    function select_option(select, value) {
        for (let i = 0; i < select.options.length; i += 1) {
            const option = select.options[i];
            if (option.__value === value) {
                option.selected = true;
                return;
            }
        }
        select.selectedIndex = -1; // no option should be selected
    }
    function select_value(select) {
        const selected_option = select.querySelector(':checked') || select.options[0];
        return selected_option && selected_option.__value;
    }
    function toggle_class(element, name, toggle) {
        element.classList[toggle ? 'add' : 'remove'](name);
    }
    function custom_event(type, detail, { bubbles = false, cancelable = false } = {}) {
        const e = document.createEvent('CustomEvent');
        e.initCustomEvent(type, bubbles, cancelable, detail);
        return e;
    }

    // we need to store the information for multiple documents because a Svelte application could also contain iframes
    // https://github.com/sveltejs/svelte/issues/3624
    const managed_styles = new Map();
    let active = 0;
    // https://github.com/darkskyapp/string-hash/blob/master/index.js
    function hash(str) {
        let hash = 5381;
        let i = str.length;
        while (i--)
            hash = ((hash << 5) - hash) ^ str.charCodeAt(i);
        return hash >>> 0;
    }
    function create_style_information(doc, node) {
        const info = { stylesheet: append_empty_stylesheet(node), rules: {} };
        managed_styles.set(doc, info);
        return info;
    }
    function create_rule(node, a, b, duration, delay, ease, fn, uid = 0) {
        const step = 16.666 / duration;
        let keyframes = '{\n';
        for (let p = 0; p <= 1; p += step) {
            const t = a + (b - a) * ease(p);
            keyframes += p * 100 + `%{${fn(t, 1 - t)}}\n`;
        }
        const rule = keyframes + `100% {${fn(b, 1 - b)}}\n}`;
        const name = `__svelte_${hash(rule)}_${uid}`;
        const doc = get_root_for_style(node);
        const { stylesheet, rules } = managed_styles.get(doc) || create_style_information(doc, node);
        if (!rules[name]) {
            rules[name] = true;
            stylesheet.insertRule(`@keyframes ${name} ${rule}`, stylesheet.cssRules.length);
        }
        const animation = node.style.animation || '';
        node.style.animation = `${animation ? `${animation}, ` : ''}${name} ${duration}ms linear ${delay}ms 1 both`;
        active += 1;
        return name;
    }
    function delete_rule(node, name) {
        const previous = (node.style.animation || '').split(', ');
        const next = previous.filter(name
            ? anim => anim.indexOf(name) < 0 // remove specific animation
            : anim => anim.indexOf('__svelte') === -1 // remove all Svelte animations
        );
        const deleted = previous.length - next.length;
        if (deleted) {
            node.style.animation = next.join(', ');
            active -= deleted;
            if (!active)
                clear_rules();
        }
    }
    function clear_rules() {
        raf(() => {
            if (active)
                return;
            managed_styles.forEach(info => {
                const { ownerNode } = info.stylesheet;
                // there is no ownerNode if it runs on jsdom.
                if (ownerNode)
                    detach(ownerNode);
            });
            managed_styles.clear();
        });
    }

    let current_component;
    function set_current_component(component) {
        current_component = component;
    }
    function get_current_component() {
        if (!current_component)
            throw new Error('Function called outside component initialization');
        return current_component;
    }
    function onMount(fn) {
        get_current_component().$$.on_mount.push(fn);
    }

    const dirty_components = [];
    const binding_callbacks = [];
    const render_callbacks = [];
    const flush_callbacks = [];
    const resolved_promise = Promise.resolve();
    let update_scheduled = false;
    function schedule_update() {
        if (!update_scheduled) {
            update_scheduled = true;
            resolved_promise.then(flush);
        }
    }
    function add_render_callback(fn) {
        render_callbacks.push(fn);
    }
    function add_flush_callback(fn) {
        flush_callbacks.push(fn);
    }
    // flush() calls callbacks in this order:
    // 1. All beforeUpdate callbacks, in order: parents before children
    // 2. All bind:this callbacks, in reverse order: children before parents.
    // 3. All afterUpdate callbacks, in order: parents before children. EXCEPT
    //    for afterUpdates called during the initial onMount, which are called in
    //    reverse order: children before parents.
    // Since callbacks might update component values, which could trigger another
    // call to flush(), the following steps guard against this:
    // 1. During beforeUpdate, any updated components will be added to the
    //    dirty_components array and will cause a reentrant call to flush(). Because
    //    the flush index is kept outside the function, the reentrant call will pick
    //    up where the earlier call left off and go through all dirty components. The
    //    current_component value is saved and restored so that the reentrant call will
    //    not interfere with the "parent" flush() call.
    // 2. bind:this callbacks cannot trigger new flush() calls.
    // 3. During afterUpdate, any updated components will NOT have their afterUpdate
    //    callback called a second time; the seen_callbacks set, outside the flush()
    //    function, guarantees this behavior.
    const seen_callbacks = new Set();
    let flushidx = 0; // Do *not* move this inside the flush() function
    function flush() {
        const saved_component = current_component;
        do {
            // first, call beforeUpdate functions
            // and update components
            while (flushidx < dirty_components.length) {
                const component = dirty_components[flushidx];
                flushidx++;
                set_current_component(component);
                update(component.$$);
            }
            set_current_component(null);
            dirty_components.length = 0;
            flushidx = 0;
            while (binding_callbacks.length)
                binding_callbacks.pop()();
            // then, once components are updated, call
            // afterUpdate functions. This may cause
            // subsequent updates...
            for (let i = 0; i < render_callbacks.length; i += 1) {
                const callback = render_callbacks[i];
                if (!seen_callbacks.has(callback)) {
                    // ...so guard against infinite loops
                    seen_callbacks.add(callback);
                    callback();
                }
            }
            render_callbacks.length = 0;
        } while (dirty_components.length);
        while (flush_callbacks.length) {
            flush_callbacks.pop()();
        }
        update_scheduled = false;
        seen_callbacks.clear();
        set_current_component(saved_component);
    }
    function update($$) {
        if ($$.fragment !== null) {
            $$.update();
            run_all($$.before_update);
            const dirty = $$.dirty;
            $$.dirty = [-1];
            $$.fragment && $$.fragment.p($$.ctx, dirty);
            $$.after_update.forEach(add_render_callback);
        }
    }

    let promise;
    function wait() {
        if (!promise) {
            promise = Promise.resolve();
            promise.then(() => {
                promise = null;
            });
        }
        return promise;
    }
    function dispatch(node, direction, kind) {
        node.dispatchEvent(custom_event(`${direction ? 'intro' : 'outro'}${kind}`));
    }
    const outroing = new Set();
    let outros;
    function group_outros() {
        outros = {
            r: 0,
            c: [],
            p: outros // parent group
        };
    }
    function check_outros() {
        if (!outros.r) {
            run_all(outros.c);
        }
        outros = outros.p;
    }
    function transition_in(block, local) {
        if (block && block.i) {
            outroing.delete(block);
            block.i(local);
        }
    }
    function transition_out(block, local, detach, callback) {
        if (block && block.o) {
            if (outroing.has(block))
                return;
            outroing.add(block);
            outros.c.push(() => {
                outroing.delete(block);
                if (callback) {
                    if (detach)
                        block.d(1);
                    callback();
                }
            });
            block.o(local);
        }
        else if (callback) {
            callback();
        }
    }
    const null_transition = { duration: 0 };
    function create_bidirectional_transition(node, fn, params, intro) {
        let config = fn(node, params);
        let t = intro ? 0 : 1;
        let running_program = null;
        let pending_program = null;
        let animation_name = null;
        function clear_animation() {
            if (animation_name)
                delete_rule(node, animation_name);
        }
        function init(program, duration) {
            const d = (program.b - t);
            duration *= Math.abs(d);
            return {
                a: t,
                b: program.b,
                d,
                duration,
                start: program.start,
                end: program.start + duration,
                group: program.group
            };
        }
        function go(b) {
            const { delay = 0, duration = 300, easing = identity, tick = noop, css } = config || null_transition;
            const program = {
                start: now() + delay,
                b
            };
            if (!b) {
                // @ts-ignore todo: improve typings
                program.group = outros;
                outros.r += 1;
            }
            if (running_program || pending_program) {
                pending_program = program;
            }
            else {
                // if this is an intro, and there's a delay, we need to do
                // an initial tick and/or apply CSS animation immediately
                if (css) {
                    clear_animation();
                    animation_name = create_rule(node, t, b, duration, delay, easing, css);
                }
                if (b)
                    tick(0, 1);
                running_program = init(program, duration);
                add_render_callback(() => dispatch(node, b, 'start'));
                loop(now => {
                    if (pending_program && now > pending_program.start) {
                        running_program = init(pending_program, duration);
                        pending_program = null;
                        dispatch(node, running_program.b, 'start');
                        if (css) {
                            clear_animation();
                            animation_name = create_rule(node, t, running_program.b, running_program.duration, 0, easing, config.css);
                        }
                    }
                    if (running_program) {
                        if (now >= running_program.end) {
                            tick(t = running_program.b, 1 - t);
                            dispatch(node, running_program.b, 'end');
                            if (!pending_program) {
                                // we're done
                                if (running_program.b) {
                                    // intro — we can tidy up immediately
                                    clear_animation();
                                }
                                else {
                                    // outro — needs to be coordinated
                                    if (!--running_program.group.r)
                                        run_all(running_program.group.c);
                                }
                            }
                            running_program = null;
                        }
                        else if (now >= running_program.start) {
                            const p = now - running_program.start;
                            t = running_program.a + running_program.d * easing(p / running_program.duration);
                            tick(t, 1 - t);
                        }
                    }
                    return !!(running_program || pending_program);
                });
            }
        }
        return {
            run(b) {
                if (is_function(config)) {
                    wait().then(() => {
                        // @ts-ignore
                        config = config();
                        go(b);
                    });
                }
                else {
                    go(b);
                }
            },
            end() {
                clear_animation();
                running_program = pending_program = null;
            }
        };
    }

    const globals = (typeof window !== 'undefined'
        ? window
        : typeof globalThis !== 'undefined'
            ? globalThis
            : global);

    function bind(component, name, callback) {
        const index = component.$$.props[name];
        if (index !== undefined) {
            component.$$.bound[index] = callback;
            callback(component.$$.ctx[index]);
        }
    }
    function create_component(block) {
        block && block.c();
    }
    function mount_component(component, target, anchor, customElement) {
        const { fragment, on_mount, on_destroy, after_update } = component.$$;
        fragment && fragment.m(target, anchor);
        if (!customElement) {
            // onMount happens before the initial afterUpdate
            add_render_callback(() => {
                const new_on_destroy = on_mount.map(run).filter(is_function);
                if (on_destroy) {
                    on_destroy.push(...new_on_destroy);
                }
                else {
                    // Edge case - component was destroyed immediately,
                    // most likely as a result of a binding initialising
                    run_all(new_on_destroy);
                }
                component.$$.on_mount = [];
            });
        }
        after_update.forEach(add_render_callback);
    }
    function destroy_component(component, detaching) {
        const $$ = component.$$;
        if ($$.fragment !== null) {
            run_all($$.on_destroy);
            $$.fragment && $$.fragment.d(detaching);
            // TODO null out other refs, including component.$$ (but need to
            // preserve final state?)
            $$.on_destroy = $$.fragment = null;
            $$.ctx = [];
        }
    }
    function make_dirty(component, i) {
        if (component.$$.dirty[0] === -1) {
            dirty_components.push(component);
            schedule_update();
            component.$$.dirty.fill(0);
        }
        component.$$.dirty[(i / 31) | 0] |= (1 << (i % 31));
    }
    function init(component, options, instance, create_fragment, not_equal, props, append_styles, dirty = [-1]) {
        const parent_component = current_component;
        set_current_component(component);
        const $$ = component.$$ = {
            fragment: null,
            ctx: null,
            // state
            props,
            update: noop,
            not_equal,
            bound: blank_object(),
            // lifecycle
            on_mount: [],
            on_destroy: [],
            on_disconnect: [],
            before_update: [],
            after_update: [],
            context: new Map(options.context || (parent_component ? parent_component.$$.context : [])),
            // everything else
            callbacks: blank_object(),
            dirty,
            skip_bound: false,
            root: options.target || parent_component.$$.root
        };
        append_styles && append_styles($$.root);
        let ready = false;
        $$.ctx = instance
            ? instance(component, options.props || {}, (i, ret, ...rest) => {
                const value = rest.length ? rest[0] : ret;
                if ($$.ctx && not_equal($$.ctx[i], $$.ctx[i] = value)) {
                    if (!$$.skip_bound && $$.bound[i])
                        $$.bound[i](value);
                    if (ready)
                        make_dirty(component, i);
                }
                return ret;
            })
            : [];
        $$.update();
        ready = true;
        run_all($$.before_update);
        // `false` as a special case of no DOM component
        $$.fragment = create_fragment ? create_fragment($$.ctx) : false;
        if (options.target) {
            if (options.hydrate) {
                const nodes = children(options.target);
                // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
                $$.fragment && $$.fragment.l(nodes);
                nodes.forEach(detach);
            }
            else {
                // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
                $$.fragment && $$.fragment.c();
            }
            if (options.intro)
                transition_in(component.$$.fragment);
            mount_component(component, options.target, options.anchor, options.customElement);
            flush();
        }
        set_current_component(parent_component);
    }
    /**
     * Base class for Svelte components. Used when dev=false.
     */
    class SvelteComponent {
        $destroy() {
            destroy_component(this, 1);
            this.$destroy = noop;
        }
        $on(type, callback) {
            const callbacks = (this.$$.callbacks[type] || (this.$$.callbacks[type] = []));
            callbacks.push(callback);
            return () => {
                const index = callbacks.indexOf(callback);
                if (index !== -1)
                    callbacks.splice(index, 1);
            };
        }
        $set($$props) {
            if (this.$$set && !is_empty($$props)) {
                this.$$.skip_bound = true;
                this.$$set($$props);
                this.$$.skip_bound = false;
            }
        }
    }

    function dispatch_dev(type, detail) {
        document.dispatchEvent(custom_event(type, Object.assign({ version: '3.50.1' }, detail), { bubbles: true }));
    }
    function append_dev(target, node) {
        dispatch_dev('SvelteDOMInsert', { target, node });
        append(target, node);
    }
    function insert_dev(target, node, anchor) {
        dispatch_dev('SvelteDOMInsert', { target, node, anchor });
        insert(target, node, anchor);
    }
    function detach_dev(node) {
        dispatch_dev('SvelteDOMRemove', { node });
        detach(node);
    }
    function listen_dev(node, event, handler, options, has_prevent_default, has_stop_propagation) {
        const modifiers = options === true ? ['capture'] : options ? Array.from(Object.keys(options)) : [];
        if (has_prevent_default)
            modifiers.push('preventDefault');
        if (has_stop_propagation)
            modifiers.push('stopPropagation');
        dispatch_dev('SvelteDOMAddEventListener', { node, event, handler, modifiers });
        const dispose = listen(node, event, handler, options);
        return () => {
            dispatch_dev('SvelteDOMRemoveEventListener', { node, event, handler, modifiers });
            dispose();
        };
    }
    function attr_dev(node, attribute, value) {
        attr(node, attribute, value);
        if (value == null)
            dispatch_dev('SvelteDOMRemoveAttribute', { node, attribute });
        else
            dispatch_dev('SvelteDOMSetAttribute', { node, attribute, value });
    }
    function set_data_dev(text, data) {
        data = '' + data;
        if (text.wholeText === data)
            return;
        dispatch_dev('SvelteDOMSetData', { node: text, data });
        text.data = data;
    }
    function validate_each_argument(arg) {
        if (typeof arg !== 'string' && !(arg && typeof arg === 'object' && 'length' in arg)) {
            let msg = '{#each} only iterates over array-like objects.';
            if (typeof Symbol === 'function' && arg && Symbol.iterator in arg) {
                msg += ' You can use a spread to convert this iterable into an array.';
            }
            throw new Error(msg);
        }
    }
    function validate_slots(name, slot, keys) {
        for (const slot_key of Object.keys(slot)) {
            if (!~keys.indexOf(slot_key)) {
                console.warn(`<${name}> received an unexpected slot "${slot_key}".`);
            }
        }
    }
    /**
     * Base class for Svelte components with some minor dev-enhancements. Used when dev=true.
     */
    class SvelteComponentDev extends SvelteComponent {
        constructor(options) {
            if (!options || (!options.target && !options.$$inline)) {
                throw new Error("'target' is a required option");
            }
            super();
        }
        $destroy() {
            super.$destroy();
            this.$destroy = () => {
                console.warn('Component was already destroyed'); // eslint-disable-line no-console
            };
        }
        $capture_state() { }
        $inject_state() { }
    }

    function cubicOut(t) {
        const f = t - 1.0;
        return f * f * f + 1.0;
    }

    function fly(node, { delay = 0, duration = 400, easing = cubicOut, x = 0, y = 0, opacity = 0 } = {}) {
        const style = getComputedStyle(node);
        const target_opacity = +style.opacity;
        const transform = style.transform === 'none' ? '' : style.transform;
        const od = target_opacity * (1 - opacity);
        return {
            delay,
            duration,
            easing,
            css: (t, u) => `
			transform: ${transform} translate(${(1 - t) * x}px, ${(1 - t) * y}px);
			opacity: ${target_opacity - (od * u)}`
        };
    }

    /* src/MultiSelect.svelte generated by Svelte v3.50.1 */

    const { Object: Object_1 } = globals;
    const file$1 = "src/MultiSelect.svelte";

    function get_each_context(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[25] = list[i];
    	return child_ctx;
    }

    function get_each_context_1(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[28] = list[i];
    	return child_ctx;
    }

    // (241:10) {#if !readonly}
    function create_if_block_2(ctx) {
    	let div;
    	let svg;
    	let path;
    	let div_title_value;

    	const block = {
    		c: function create() {
    			div = element("div");
    			svg = svg_element("svg");
    			path = svg_element("path");
    			attr_dev(path, "d", iconClearPath);
    			attr_dev(path, "class", "svelte-1i3qjsm");
    			add_location(path, file$1, 243, 16, 6528);
    			attr_dev(svg, "class", "icon-clear svelte-1i3qjsm");
    			attr_dev(svg, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg, "width", "18");
    			attr_dev(svg, "height", "18");
    			attr_dev(svg, "viewBox", "0 0 24 24");
    			add_location(svg, file$1, 242, 14, 6409);
    			attr_dev(div, "class", "token-remove svelte-1i3qjsm");
    			attr_dev(div, "title", div_title_value = "Remove " + /*s*/ ctx[28].name);
    			add_location(div, file$1, 241, 12, 6344);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, svg);
    			append_dev(svg, path);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*selected*/ 32 && div_title_value !== (div_title_value = "Remove " + /*s*/ ctx[28].name)) {
    				attr_dev(div, "title", div_title_value);
    			}
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_2.name,
    		type: "if",
    		source: "(241:10) {#if !readonly}",
    		ctx
    	});

    	return block;
    }

    // (238:6) {#each Object.values(selected) as s}
    function create_each_block_1(ctx) {
    	let div;
    	let span;
    	let t0_value = /*s*/ ctx[28].name + "";
    	let t0;
    	let t1;
    	let div_data_id_value;
    	let if_block = !/*readonly*/ ctx[1] && create_if_block_2(ctx);

    	const block = {
    		c: function create() {
    			div = element("div");
    			span = element("span");
    			t0 = text(t0_value);
    			t1 = space();
    			if (if_block) if_block.c();
    			add_location(span, file$1, 239, 10, 6284);
    			attr_dev(div, "class", "token svelte-1i3qjsm");
    			attr_dev(div, "data-id", div_data_id_value = /*s*/ ctx[28].value);
    			add_location(div, file$1, 238, 8, 6234);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, span);
    			append_dev(span, t0);
    			append_dev(div, t1);
    			if (if_block) if_block.m(div, null);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*selected*/ 32 && t0_value !== (t0_value = /*s*/ ctx[28].name + "")) set_data_dev(t0, t0_value);

    			if (!/*readonly*/ ctx[1]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);
    				} else {
    					if_block = create_if_block_2(ctx);
    					if_block.c();
    					if_block.m(div, null);
    				}
    			} else if (if_block) {
    				if_block.d(1);
    				if_block = null;
    			}

    			if (dirty & /*selected*/ 32 && div_data_id_value !== (div_data_id_value = /*s*/ ctx[28].value)) {
    				attr_dev(div, "data-id", div_data_id_value);
    			}
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			if (if_block) if_block.d();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block_1.name,
    		type: "each",
    		source: "(238:6) {#each Object.values(selected) as s}",
    		ctx
    	});

    	return block;
    }

    // (251:8) {#if !readonly}
    function create_if_block_1(ctx) {
    	let input_1;
    	let t0;
    	let div;
    	let svg0;
    	let path0;
    	let t1;
    	let svg1;
    	let path1;
    	let mounted;
    	let dispose;

    	const block = {
    		c: function create() {
    			input_1 = element("input");
    			t0 = space();
    			div = element("div");
    			svg0 = svg_element("svg");
    			path0 = svg_element("path");
    			t1 = space();
    			svg1 = svg_element("svg");
    			path1 = svg_element("path");
    			attr_dev(input_1, "id", /*id*/ ctx[0]);
    			attr_dev(input_1, "autocomplete", "off");
    			attr_dev(input_1, "placeholder", /*placeholder*/ ctx[2]);
    			attr_dev(input_1, "class", "svelte-1i3qjsm");
    			add_location(input_1, file$1, 251, 10, 6703);
    			attr_dev(path0, "d", iconClearPath);
    			attr_dev(path0, "class", "svelte-1i3qjsm");
    			add_location(path0, file$1, 254, 14, 7079);
    			attr_dev(svg0, "class", "icon-clear svelte-1i3qjsm");
    			attr_dev(svg0, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg0, "width", "18");
    			attr_dev(svg0, "height", "18");
    			attr_dev(svg0, "viewBox", "0 0 24 24");
    			add_location(svg0, file$1, 253, 12, 6962);
    			attr_dev(div, "class", "remove-all svelte-1i3qjsm");
    			attr_dev(div, "title", "Remove All");
    			toggle_class(div, "hidden", !Object.keys(/*selected*/ ctx[5]).length);
    			add_location(div, file$1, 252, 10, 6861);
    			attr_dev(path1, "d", "M5 8l4 4 4-4z");
    			attr_dev(path1, "class", "svelte-1i3qjsm");
    			add_location(path1, file$1, 257, 116, 7259);
    			attr_dev(svg1, "class", "dropdown-arrow svelte-1i3qjsm");
    			attr_dev(svg1, "xmlns", "http://www.w3.org/2000/svg");
    			attr_dev(svg1, "width", "18");
    			attr_dev(svg1, "height", "18");
    			attr_dev(svg1, "viewBox", "0 0 18 18");
    			add_location(svg1, file$1, 257, 10, 7153);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, input_1, anchor);
    			set_input_value(input_1, /*inputValue*/ ctx[3]);
    			/*input_1_binding*/ ctx[20](input_1);
    			insert_dev(target, t0, anchor);
    			insert_dev(target, div, anchor);
    			append_dev(div, svg0);
    			append_dev(svg0, path0);
    			insert_dev(target, t1, anchor);
    			insert_dev(target, svg1, anchor);
    			append_dev(svg1, path1);

    			if (!mounted) {
    				dispose = [
    					listen_dev(input_1, "input", /*input_1_input_handler*/ ctx[19]),
    					listen_dev(input_1, "keyup", /*handleKeyup*/ ctx[10], false, false, false),
    					listen_dev(input_1, "blur", /*handleBlur*/ ctx[11], false, false, false)
    				];

    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*id*/ 1) {
    				attr_dev(input_1, "id", /*id*/ ctx[0]);
    			}

    			if (dirty & /*placeholder*/ 4) {
    				attr_dev(input_1, "placeholder", /*placeholder*/ ctx[2]);
    			}

    			if (dirty & /*inputValue*/ 8 && input_1.value !== /*inputValue*/ ctx[3]) {
    				set_input_value(input_1, /*inputValue*/ ctx[3]);
    			}

    			if (dirty & /*Object, selected*/ 32) {
    				toggle_class(div, "hidden", !Object.keys(/*selected*/ ctx[5]).length);
    			}
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(input_1);
    			/*input_1_binding*/ ctx[20](null);
    			if (detaching) detach_dev(t0);
    			if (detaching) detach_dev(div);
    			if (detaching) detach_dev(t1);
    			if (detaching) detach_dev(svg1);
    			mounted = false;
    			run_all(dispose);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block_1.name,
    		type: "if",
    		source: "(251:8) {#if !readonly}",
    		ctx
    	});

    	return block;
    }

    // (265:4) {#if showOptions}
    function create_if_block(ctx) {
    	let ul;
    	let ul_transition;
    	let current;
    	let mounted;
    	let dispose;
    	let each_value = /*filtered*/ ctx[6];
    	validate_each_argument(each_value);
    	let each_blocks = [];

    	for (let i = 0; i < each_value.length; i += 1) {
    		each_blocks[i] = create_each_block(get_each_context(ctx, each_value, i));
    	}

    	const block = {
    		c: function create() {
    			ul = element("ul");

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			attr_dev(ul, "class", "options svelte-1i3qjsm");
    			add_location(ul, file$1, 265, 6, 7454);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, ul, anchor);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].m(ul, null);
    			}

    			current = true;

    			if (!mounted) {
    				dispose = listen_dev(ul, "mousedown", prevent_default(/*handleOptionMousedown*/ ctx[13]), false, true, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*filtered, selected, activeOption*/ 112) {
    				each_value = /*filtered*/ ctx[6];
    				validate_each_argument(each_value);
    				let i;

    				for (i = 0; i < each_value.length; i += 1) {
    					const child_ctx = get_each_context(ctx, each_value, i);

    					if (each_blocks[i]) {
    						each_blocks[i].p(child_ctx, dirty);
    					} else {
    						each_blocks[i] = create_each_block(child_ctx);
    						each_blocks[i].c();
    						each_blocks[i].m(ul, null);
    					}
    				}

    				for (; i < each_blocks.length; i += 1) {
    					each_blocks[i].d(1);
    				}

    				each_blocks.length = each_value.length;
    			}
    		},
    		i: function intro(local) {
    			if (current) return;

    			add_render_callback(() => {
    				if (!ul_transition) ul_transition = create_bidirectional_transition(ul, fly, { duration: 200, y: 5 }, true);
    				ul_transition.run(1);
    			});

    			current = true;
    		},
    		o: function outro(local) {
    			if (!ul_transition) ul_transition = create_bidirectional_transition(ul, fly, { duration: 200, y: 5 }, false);
    			ul_transition.run(0);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(ul);
    			destroy_each(each_blocks, detaching);
    			if (detaching && ul_transition) ul_transition.end();
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block.name,
    		type: "if",
    		source: "(265:4) {#if showOptions}",
    		ctx
    	});

    	return block;
    }

    // (267:8) {#each filtered as option}
    function create_each_block(ctx) {
    	let li;
    	let t_value = /*option*/ ctx[25].name + "";
    	let t;
    	let li_data_value_value;

    	const block = {
    		c: function create() {
    			li = element("li");
    			t = text(t_value);
    			attr_dev(li, "data-value", li_data_value_value = /*option*/ ctx[25].value);
    			attr_dev(li, "class", "svelte-1i3qjsm");
    			toggle_class(li, "selected", /*selected*/ ctx[5][/*option*/ ctx[25].value]);
    			toggle_class(li, "active", /*activeOption*/ ctx[4] === /*option*/ ctx[25]);
    			add_location(li, file$1, 267, 10, 7613);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, li, anchor);
    			append_dev(li, t);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*filtered*/ 64 && t_value !== (t_value = /*option*/ ctx[25].name + "")) set_data_dev(t, t_value);

    			if (dirty & /*filtered*/ 64 && li_data_value_value !== (li_data_value_value = /*option*/ ctx[25].value)) {
    				attr_dev(li, "data-value", li_data_value_value);
    			}

    			if (dirty & /*selected, filtered*/ 96) {
    				toggle_class(li, "selected", /*selected*/ ctx[5][/*option*/ ctx[25].value]);
    			}

    			if (dirty & /*activeOption, filtered*/ 80) {
    				toggle_class(li, "active", /*activeOption*/ ctx[4] === /*option*/ ctx[25]);
    			}
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(li);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block.name,
    		type: "each",
    		source: "(267:8) {#each filtered as option}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$1(ctx) {
    	let div2;
    	let div1;
    	let t0;
    	let div0;
    	let t1;
    	let select;
    	let t2;
    	let current;
    	let mounted;
    	let dispose;
    	let each_value_1 = Object.values(/*selected*/ ctx[5]);
    	validate_each_argument(each_value_1);
    	let each_blocks = [];

    	for (let i = 0; i < each_value_1.length; i += 1) {
    		each_blocks[i] = create_each_block_1(get_each_context_1(ctx, each_value_1, i));
    	}

    	let if_block0 = !/*readonly*/ ctx[1] && create_if_block_1(ctx);
    	const default_slot_template = /*#slots*/ ctx[18].default;
    	const default_slot = create_slot(default_slot_template, ctx, /*$$scope*/ ctx[17], null);
    	let if_block1 = /*showOptions*/ ctx[8] && create_if_block(ctx);

    	const block = {
    		c: function create() {
    			div2 = element("div");
    			div1 = element("div");

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			t0 = space();
    			div0 = element("div");
    			if (if_block0) if_block0.c();
    			t1 = space();
    			select = element("select");
    			if (default_slot) default_slot.c();
    			t2 = space();
    			if (if_block1) if_block1.c();
    			attr_dev(div0, "class", "actions svelte-1i3qjsm");
    			add_location(div0, file$1, 249, 6, 6647);
    			attr_dev(div1, "class", "tokens svelte-1i3qjsm");
    			toggle_class(div1, "showOptions", /*showOptions*/ ctx[8]);
    			add_location(div1, file$1, 236, 4, 6116);
    			attr_dev(select, "type", "multiple");
    			attr_dev(select, "class", "hidden svelte-1i3qjsm");
    			add_location(select, file$1, 262, 4, 7342);
    			attr_dev(div2, "class", "multiselect svelte-1i3qjsm");
    			toggle_class(div2, "readonly", /*readonly*/ ctx[1]);
    			add_location(div2, file$1, 235, 2, 6071);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div2, anchor);
    			append_dev(div2, div1);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].m(div1, null);
    			}

    			append_dev(div1, t0);
    			append_dev(div1, div0);
    			if (if_block0) if_block0.m(div0, null);
    			append_dev(div2, t1);
    			append_dev(div2, select);

    			if (default_slot) {
    				default_slot.m(select, null);
    			}

    			/*select_binding*/ ctx[21](select);
    			append_dev(div2, t2);
    			if (if_block1) if_block1.m(div2, null);
    			current = true;

    			if (!mounted) {
    				dispose = listen_dev(div1, "click", /*handleTokenClick*/ ctx[12], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*Object, selected, iconClearPath, readonly*/ 34) {
    				each_value_1 = Object.values(/*selected*/ ctx[5]);
    				validate_each_argument(each_value_1);
    				let i;

    				for (i = 0; i < each_value_1.length; i += 1) {
    					const child_ctx = get_each_context_1(ctx, each_value_1, i);

    					if (each_blocks[i]) {
    						each_blocks[i].p(child_ctx, dirty);
    					} else {
    						each_blocks[i] = create_each_block_1(child_ctx);
    						each_blocks[i].c();
    						each_blocks[i].m(div1, t0);
    					}
    				}

    				for (; i < each_blocks.length; i += 1) {
    					each_blocks[i].d(1);
    				}

    				each_blocks.length = each_value_1.length;
    			}

    			if (!/*readonly*/ ctx[1]) {
    				if (if_block0) {
    					if_block0.p(ctx, dirty);
    				} else {
    					if_block0 = create_if_block_1(ctx);
    					if_block0.c();
    					if_block0.m(div0, null);
    				}
    			} else if (if_block0) {
    				if_block0.d(1);
    				if_block0 = null;
    			}

    			if (!current || dirty & /*showOptions*/ 256) {
    				toggle_class(div1, "showOptions", /*showOptions*/ ctx[8]);
    			}

    			if (default_slot) {
    				if (default_slot.p && (!current || dirty & /*$$scope*/ 131072)) {
    					update_slot_base(
    						default_slot,
    						default_slot_template,
    						ctx,
    						/*$$scope*/ ctx[17],
    						!current
    						? get_all_dirty_from_scope(/*$$scope*/ ctx[17])
    						: get_slot_changes(default_slot_template, /*$$scope*/ ctx[17], dirty, null),
    						null
    					);
    				}
    			}

    			if (/*showOptions*/ ctx[8]) {
    				if (if_block1) {
    					if_block1.p(ctx, dirty);

    					if (dirty & /*showOptions*/ 256) {
    						transition_in(if_block1, 1);
    					}
    				} else {
    					if_block1 = create_if_block(ctx);
    					if_block1.c();
    					transition_in(if_block1, 1);
    					if_block1.m(div2, null);
    				}
    			} else if (if_block1) {
    				group_outros();

    				transition_out(if_block1, 1, 1, () => {
    					if_block1 = null;
    				});

    				check_outros();
    			}

    			if (!current || dirty & /*readonly*/ 2) {
    				toggle_class(div2, "readonly", /*readonly*/ ctx[1]);
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(default_slot, local);
    			transition_in(if_block1);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(default_slot, local);
    			transition_out(if_block1);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div2);
    			destroy_each(each_blocks, detaching);
    			if (if_block0) if_block0.d();
    			if (default_slot) default_slot.d(detaching);
    			/*select_binding*/ ctx[21](null);
    			if (if_block1) if_block1.d();
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$1.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    const iconClearPath = 'M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z';

    function instance$1($$self, $$props, $$invalidate) {
    	let filtered;
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('MultiSelect', slots, ['default']);
    	let { id = '' } = $$props;
    	let { value = [] } = $$props;
    	let { readonly = false } = $$props;
    	let { placeholder = '' } = $$props;

    	let input,
    		inputValue,
    		options = [],
    		activeOption,
    		showOptions = false,
    		selected = {},
    		first = true,
    		slot;

    	onMount(() => {
    		slot.querySelectorAll('option').forEach(o => {
    			o.selected && !value.includes(o.value) && $$invalidate(14, value = [...value, o.value]);
    			$$invalidate(15, options = [...options, { value: o.value, name: o.textContent }]);
    		});

    		value && $$invalidate(5, selected = options.reduce(
    			(obj, op) => value.includes(op.value)
    			? { ...obj, [op.value]: op }
    			: obj,
    			{}
    		));

    		$$invalidate(16, first = false);
    	});

    	function add(token) {
    		if (!readonly) $$invalidate(5, selected[token.value] = token, selected);
    	}

    	function remove(value) {
    		if (!readonly) {
    			const { [value]: val, ...rest } = selected;
    			$$invalidate(5, selected = rest);
    		}
    	}

    	function optionsVisibility(show) {
    		if (readonly) return;

    		if (typeof show === 'boolean') {
    			$$invalidate(8, showOptions = show);
    			show && input.focus();
    		} else {
    			$$invalidate(8, showOptions = !showOptions);
    		}

    		if (!showOptions) {
    			$$invalidate(4, activeOption = undefined);
    		}
    	}

    	function handleKeyup(e) {
    		if (e.keyCode === 13) {
    			Object.keys(selected).includes(activeOption.value)
    			? remove(activeOption.value)
    			: add(activeOption);

    			$$invalidate(3, inputValue = '');
    		}

    		if ([38, 40].includes(e.keyCode)) {
    			// up and down arrows
    			const increment = e.keyCode === 38 ? -1 : 1;

    			const calcIndex = filtered.indexOf(activeOption) + increment;

    			$$invalidate(4, activeOption = calcIndex < 0
    			? filtered[filtered.length - 1]
    			: calcIndex === filtered.length
    				? filtered[0]
    				: filtered[calcIndex]);
    		}
    	}

    	function handleBlur(e) {
    		optionsVisibility(false);
    	}

    	function handleTokenClick(e) {
    		if (e.target.closest('.token-remove')) {
    			e.stopPropagation();
    			remove(e.target.closest('.token').dataset.id);
    		} else if (e.target.closest('.remove-all')) {
    			$$invalidate(5, selected = []);
    			$$invalidate(3, inputValue = '');
    		} else {
    			optionsVisibility(true);
    		}
    	}

    	function handleOptionMousedown(e) {
    		const value = e.target.dataset.value;

    		if (selected[value]) {
    			remove(value);
    		} else {
    			add(options.filter(o => o.value === value)[0]);
    			input.focus();
    		}
    	}

    	const writable_props = ['id', 'value', 'readonly', 'placeholder'];

    	Object_1.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console.warn(`<MultiSelect> was created with unknown prop '${key}'`);
    	});

    	function input_1_input_handler() {
    		inputValue = this.value;
    		$$invalidate(3, inputValue);
    	}

    	function input_1_binding($$value) {
    		binding_callbacks[$$value ? 'unshift' : 'push'](() => {
    			input = $$value;
    			$$invalidate(7, input);
    		});
    	}

    	function select_binding($$value) {
    		binding_callbacks[$$value ? 'unshift' : 'push'](() => {
    			slot = $$value;
    			$$invalidate(9, slot);
    		});
    	}

    	$$self.$$set = $$props => {
    		if ('id' in $$props) $$invalidate(0, id = $$props.id);
    		if ('value' in $$props) $$invalidate(14, value = $$props.value);
    		if ('readonly' in $$props) $$invalidate(1, readonly = $$props.readonly);
    		if ('placeholder' in $$props) $$invalidate(2, placeholder = $$props.placeholder);
    		if ('$$scope' in $$props) $$invalidate(17, $$scope = $$props.$$scope);
    	};

    	$$self.$capture_state = () => ({
    		onMount,
    		fly,
    		id,
    		value,
    		readonly,
    		placeholder,
    		input,
    		inputValue,
    		options,
    		activeOption,
    		showOptions,
    		selected,
    		first,
    		slot,
    		iconClearPath,
    		add,
    		remove,
    		optionsVisibility,
    		handleKeyup,
    		handleBlur,
    		handleTokenClick,
    		handleOptionMousedown,
    		filtered
    	});

    	$$self.$inject_state = $$props => {
    		if ('id' in $$props) $$invalidate(0, id = $$props.id);
    		if ('value' in $$props) $$invalidate(14, value = $$props.value);
    		if ('readonly' in $$props) $$invalidate(1, readonly = $$props.readonly);
    		if ('placeholder' in $$props) $$invalidate(2, placeholder = $$props.placeholder);
    		if ('input' in $$props) $$invalidate(7, input = $$props.input);
    		if ('inputValue' in $$props) $$invalidate(3, inputValue = $$props.inputValue);
    		if ('options' in $$props) $$invalidate(15, options = $$props.options);
    		if ('activeOption' in $$props) $$invalidate(4, activeOption = $$props.activeOption);
    		if ('showOptions' in $$props) $$invalidate(8, showOptions = $$props.showOptions);
    		if ('selected' in $$props) $$invalidate(5, selected = $$props.selected);
    		if ('first' in $$props) $$invalidate(16, first = $$props.first);
    		if ('slot' in $$props) $$invalidate(9, slot = $$props.slot);
    		if ('filtered' in $$props) $$invalidate(6, filtered = $$props.filtered);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	$$self.$$.update = () => {
    		if ($$self.$$.dirty & /*first, selected*/ 65568) {
    			if (!first) $$invalidate(14, value = Object.values(selected).map(o => o.value));
    		}

    		if ($$self.$$.dirty & /*options, inputValue*/ 32776) {
    			$$invalidate(6, filtered = options.filter(o => inputValue
    			? o.name.toLowerCase().includes(inputValue.toLowerCase())
    			: o));
    		}

    		if ($$self.$$.dirty & /*activeOption, filtered, inputValue*/ 88) {
    			if (activeOption && !filtered.includes(activeOption) || !activeOption && inputValue) $$invalidate(4, activeOption = filtered[0]);
    		}
    	};

    	return [
    		id,
    		readonly,
    		placeholder,
    		inputValue,
    		activeOption,
    		selected,
    		filtered,
    		input,
    		showOptions,
    		slot,
    		handleKeyup,
    		handleBlur,
    		handleTokenClick,
    		handleOptionMousedown,
    		value,
    		options,
    		first,
    		$$scope,
    		slots,
    		input_1_input_handler,
    		input_1_binding,
    		select_binding
    	];
    }

    class MultiSelect extends SvelteComponentDev {
    	constructor(options) {
    		super(options);

    		init(this, options, instance$1, create_fragment$1, safe_not_equal, {
    			id: 0,
    			value: 14,
    			readonly: 1,
    			placeholder: 2
    		});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "MultiSelect",
    			options,
    			id: create_fragment$1.name
    		});
    	}

    	get id() {
    		throw new Error("<MultiSelect>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set id(value) {
    		throw new Error("<MultiSelect>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get value() {
    		throw new Error("<MultiSelect>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set value(value) {
    		throw new Error("<MultiSelect>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get readonly() {
    		throw new Error("<MultiSelect>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set readonly(value) {
    		throw new Error("<MultiSelect>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get placeholder() {
    		throw new Error("<MultiSelect>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set placeholder(value) {
    		throw new Error("<MultiSelect>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    /* src/AdminUpdateUser.svelte generated by Svelte v3.50.1 */

    const { console: console_1 } = globals;
    const file = "src/AdminUpdateUser.svelte";

    // (55:0) <MultiSelect bind:value={user_group} placeholder="User Group(s)">
    function create_default_slot(ctx) {
    	let option0;
    	let t1;
    	let option1;
    	let t3;
    	let option2;

    	const block = {
    		c: function create() {
    			option0 = element("option");
    			option0.textContent = "Project Lead";
    			t1 = space();
    			option1 = element("option");
    			option1.textContent = "Project Manager";
    			t3 = space();
    			option2 = element("option");
    			option2.textContent = "Team Member";
    			option0.__value = "Project Lead";
    			option0.value = option0.__value;
    			add_location(option0, file, 55, 1, 1393);
    			option1.__value = "Project Manager";
    			option1.value = option1.__value;
    			add_location(option1, file, 56, 1, 1446);
    			option2.__value = "Team Member";
    			option2.value = option2.__value;
    			add_location(option2, file, 57, 1, 1504);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, option0, anchor);
    			insert_dev(target, t1, anchor);
    			insert_dev(target, option1, anchor);
    			insert_dev(target, t3, anchor);
    			insert_dev(target, option2, anchor);
    		},
    		p: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(option0);
    			if (detaching) detach_dev(t1);
    			if (detaching) detach_dev(option1);
    			if (detaching) detach_dev(t3);
    			if (detaching) detach_dev(option2);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_default_slot.name,
    		type: "slot",
    		source: "(55:0) <MultiSelect bind:value={user_group} placeholder=\\\"User Group(s)\\\">",
    		ctx
    	});

    	return block;
    }

    function create_fragment(ctx) {
    	let div;
    	let h1;
    	let t1;
    	let input0;
    	let t2;
    	let br0;
    	let t3;
    	let input1;
    	let t4;
    	let br1;
    	let t5;
    	let input2;
    	let t6;
    	let br2;
    	let t7;
    	let multiselect;
    	let updating_value;
    	let t8;
    	let br3;
    	let t9;
    	let select;
    	let option0;
    	let option1;
    	let t12;
    	let br4;
    	let t13;
    	let button;
    	let t15;
    	let p0;
    	let t16;
    	let t17;
    	let t18;
    	let p1;
    	let t19;
    	let t20;
    	let t21;
    	let ul;
    	let li0;
    	let t22;
    	let t23;
    	let t24;
    	let li1;
    	let t25;
    	let t26;
    	let t27;
    	let li2;
    	let t28;
    	let t29;
    	let t30;
    	let li3;
    	let t31;
    	let t32;
    	let t33;
    	let li4;
    	let t34;
    	let t35;
    	let current;
    	let mounted;
    	let dispose;

    	function multiselect_value_binding(value) {
    		/*multiselect_value_binding*/ ctx[11](value);
    	}

    	let multiselect_props = {
    		placeholder: "User Group(s)",
    		$$slots: { default: [create_default_slot] },
    		$$scope: { ctx }
    	};

    	if (/*user_group*/ ctx[3] !== void 0) {
    		multiselect_props.value = /*user_group*/ ctx[3];
    	}

    	multiselect = new MultiSelect({ props: multiselect_props, $$inline: true });
    	binding_callbacks.push(() => bind(multiselect, 'value', multiselect_value_binding));

    	const block = {
    		c: function create() {
    			div = element("div");
    			h1 = element("h1");
    			h1.textContent = "Admin Update User";
    			t1 = space();
    			input0 = element("input");
    			t2 = space();
    			br0 = element("br");
    			t3 = space();
    			input1 = element("input");
    			t4 = space();
    			br1 = element("br");
    			t5 = space();
    			input2 = element("input");
    			t6 = space();
    			br2 = element("br");
    			t7 = space();
    			create_component(multiselect.$$.fragment);
    			t8 = space();
    			br3 = element("br");
    			t9 = space();
    			select = element("select");
    			option0 = element("option");
    			option0.textContent = "Active";
    			option1 = element("option");
    			option1.textContent = "Inactive";
    			t12 = space();
    			br4 = element("br");
    			t13 = space();
    			button = element("button");
    			button.textContent = "Update User";
    			t15 = space();
    			p0 = element("p");
    			t16 = text("Message: ");
    			t17 = text(/*result*/ ctx[5]);
    			t18 = space();
    			p1 = element("p");
    			t19 = text("Code: ");
    			t20 = text(/*code*/ ctx[6]);
    			t21 = space();
    			ul = element("ul");
    			li0 = element("li");
    			t22 = text("Username: ");
    			t23 = text(/*username*/ ctx[0]);
    			t24 = space();
    			li1 = element("li");
    			t25 = text("Password: ");
    			t26 = text(/*password*/ ctx[1]);
    			t27 = space();
    			li2 = element("li");
    			t28 = text("Email: ");
    			t29 = text(/*email*/ ctx[2]);
    			t30 = space();
    			li3 = element("li");
    			t31 = text("UserGroup: ");
    			t32 = text(/*user_group*/ ctx[3]);
    			t33 = space();
    			li4 = element("li");
    			t34 = text("Status: ");
    			t35 = text(/*status*/ ctx[4]);
    			attr_dev(h1, "class", "svelte-i7qo5m");
    			add_location(h1, file, 49, 0, 1003);
    			attr_dev(input0, "type", "text");
    			attr_dev(input0, "placeholder", "Username");
    			add_location(input0, file, 50, 0, 1030);
    			add_location(br0, file, 50, 66, 1096);
    			attr_dev(input1, "type", "password");
    			attr_dev(input1, "placeholder", "Password");
    			add_location(input1, file, 51, 0, 1101);
    			add_location(br1, file, 51, 70, 1171);
    			attr_dev(input2, "type", "email");
    			attr_dev(input2, "placeholder", "Email");
    			add_location(input2, file, 52, 0, 1176);
    			add_location(br2, file, 52, 61, 1237);
    			add_location(br3, file, 58, 15, 1568);
    			option0.__value = "Active";
    			option0.value = option0.__value;
    			add_location(option0, file, 60, 1, 1624);
    			option1.__value = "Inactive";
    			option1.value = option1.__value;
    			add_location(option1, file, 61, 1, 1665);
    			attr_dev(select, "placeholder", "Status");
    			if (/*status*/ ctx[4] === void 0) add_render_callback(() => /*select_change_handler*/ ctx[12].call(select));
    			add_location(select, file, 59, 0, 1573);
    			add_location(br4, file, 62, 10, 1718);
    			add_location(button, file, 64, 0, 1724);
    			set_style(div, "text-align", "center");
    			add_location(div, file, 48, 0, 971);
    			add_location(p0, file, 67, 0, 1786);
    			add_location(p1, file, 68, 0, 1811);
    			add_location(li0, file, 71, 1, 1838);
    			add_location(li1, file, 72, 1, 1869);
    			add_location(li2, file, 73, 1, 1900);
    			add_location(li3, file, 74, 1, 1925);
    			add_location(li4, file, 75, 1, 1959);
    			add_location(ul, file, 70, 0, 1832);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, h1);
    			append_dev(div, t1);
    			append_dev(div, input0);
    			set_input_value(input0, /*username*/ ctx[0]);
    			append_dev(div, t2);
    			append_dev(div, br0);
    			append_dev(div, t3);
    			append_dev(div, input1);
    			set_input_value(input1, /*password*/ ctx[1]);
    			append_dev(div, t4);
    			append_dev(div, br1);
    			append_dev(div, t5);
    			append_dev(div, input2);
    			set_input_value(input2, /*email*/ ctx[2]);
    			append_dev(div, t6);
    			append_dev(div, br2);
    			append_dev(div, t7);
    			mount_component(multiselect, div, null);
    			append_dev(div, t8);
    			append_dev(div, br3);
    			append_dev(div, t9);
    			append_dev(div, select);
    			append_dev(select, option0);
    			append_dev(select, option1);
    			select_option(select, /*status*/ ctx[4]);
    			append_dev(div, t12);
    			append_dev(div, br4);
    			append_dev(div, t13);
    			append_dev(div, button);
    			insert_dev(target, t15, anchor);
    			insert_dev(target, p0, anchor);
    			append_dev(p0, t16);
    			append_dev(p0, t17);
    			insert_dev(target, t18, anchor);
    			insert_dev(target, p1, anchor);
    			append_dev(p1, t19);
    			append_dev(p1, t20);
    			insert_dev(target, t21, anchor);
    			insert_dev(target, ul, anchor);
    			append_dev(ul, li0);
    			append_dev(li0, t22);
    			append_dev(li0, t23);
    			append_dev(ul, t24);
    			append_dev(ul, li1);
    			append_dev(li1, t25);
    			append_dev(li1, t26);
    			append_dev(ul, t27);
    			append_dev(ul, li2);
    			append_dev(li2, t28);
    			append_dev(li2, t29);
    			append_dev(ul, t30);
    			append_dev(ul, li3);
    			append_dev(li3, t31);
    			append_dev(li3, t32);
    			append_dev(ul, t33);
    			append_dev(ul, li4);
    			append_dev(li4, t34);
    			append_dev(li4, t35);
    			current = true;

    			if (!mounted) {
    				dispose = [
    					listen_dev(input0, "input", /*input0_input_handler*/ ctx[8]),
    					listen_dev(input1, "input", /*input1_input_handler*/ ctx[9]),
    					listen_dev(input2, "input", /*input2_input_handler*/ ctx[10]),
    					listen_dev(select, "change", /*select_change_handler*/ ctx[12]),
    					listen_dev(button, "click", /*handleClick*/ ctx[7], false, false, false)
    				];

    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*username*/ 1 && input0.value !== /*username*/ ctx[0]) {
    				set_input_value(input0, /*username*/ ctx[0]);
    			}

    			if (dirty & /*password*/ 2 && input1.value !== /*password*/ ctx[1]) {
    				set_input_value(input1, /*password*/ ctx[1]);
    			}

    			if (dirty & /*email*/ 4 && input2.value !== /*email*/ ctx[2]) {
    				set_input_value(input2, /*email*/ ctx[2]);
    			}

    			const multiselect_changes = {};

    			if (dirty & /*$$scope*/ 16384) {
    				multiselect_changes.$$scope = { dirty, ctx };
    			}

    			if (!updating_value && dirty & /*user_group*/ 8) {
    				updating_value = true;
    				multiselect_changes.value = /*user_group*/ ctx[3];
    				add_flush_callback(() => updating_value = false);
    			}

    			multiselect.$set(multiselect_changes);

    			if (dirty & /*status*/ 16) {
    				select_option(select, /*status*/ ctx[4]);
    			}

    			if (!current || dirty & /*result*/ 32) set_data_dev(t17, /*result*/ ctx[5]);
    			if (!current || dirty & /*code*/ 64) set_data_dev(t20, /*code*/ ctx[6]);
    			if (!current || dirty & /*username*/ 1) set_data_dev(t23, /*username*/ ctx[0]);
    			if (!current || dirty & /*password*/ 2) set_data_dev(t26, /*password*/ ctx[1]);
    			if (!current || dirty & /*email*/ 4) set_data_dev(t29, /*email*/ ctx[2]);
    			if (!current || dirty & /*user_group*/ 8) set_data_dev(t32, /*user_group*/ ctx[3]);
    			if (!current || dirty & /*status*/ 16) set_data_dev(t35, /*status*/ ctx[4]);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(multiselect.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(multiselect.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			destroy_component(multiselect);
    			if (detaching) detach_dev(t15);
    			if (detaching) detach_dev(p0);
    			if (detaching) detach_dev(t18);
    			if (detaching) detach_dev(p1);
    			if (detaching) detach_dev(t21);
    			if (detaching) detach_dev(ul);
    			mounted = false;
    			run_all(dispose);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots('AdminUpdateUser', slots, []);
    	let { username = "" } = $$props;
    	let { password = "" } = $$props;
    	let { email = "" } = $$props;
    	let { user_group = "" } = $$props;
    	let { status = "" } = $$props;
    	let { result = null } = $$props;
    	let { code = null } = $$props;
    	let value;

    	async function handleClick() {
    		$$invalidate(3, user_group = user_group.toString());

    		await fetch("http://localhost:4000/admin-update-user", {
    			mode: "cors",
    			method: "POST",
    			cache: "no-cache",
    			headers: {
    				"Content-Type": "application/json",
    				"Access-Control-Allow-Origin": "*"
    			},
    			body: JSON.stringify({
    				username,
    				password,
    				email,
    				user_group,
    				status
    			})
    		}).then(async response => {
    			if (response) {
    				// console.log(response)
    				// result = response.status
    				const json = await response.json();

    				// result = JSON.stringify(json)
    				$$invalidate(5, result = json.error);

    				$$invalidate(6, code = json.code);
    			}
    		}).catch(error => {
    			console.error(error);
    		});
    	}

    	const writable_props = ['username', 'password', 'email', 'user_group', 'status', 'result', 'code'];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== '$$' && key !== 'slot') console_1.warn(`<AdminUpdateUser> was created with unknown prop '${key}'`);
    	});

    	function input0_input_handler() {
    		username = this.value;
    		$$invalidate(0, username);
    	}

    	function input1_input_handler() {
    		password = this.value;
    		$$invalidate(1, password);
    	}

    	function input2_input_handler() {
    		email = this.value;
    		$$invalidate(2, email);
    	}

    	function multiselect_value_binding(value) {
    		user_group = value;
    		$$invalidate(3, user_group);
    	}

    	function select_change_handler() {
    		status = select_value(this);
    		$$invalidate(4, status);
    	}

    	$$self.$$set = $$props => {
    		if ('username' in $$props) $$invalidate(0, username = $$props.username);
    		if ('password' in $$props) $$invalidate(1, password = $$props.password);
    		if ('email' in $$props) $$invalidate(2, email = $$props.email);
    		if ('user_group' in $$props) $$invalidate(3, user_group = $$props.user_group);
    		if ('status' in $$props) $$invalidate(4, status = $$props.status);
    		if ('result' in $$props) $$invalidate(5, result = $$props.result);
    		if ('code' in $$props) $$invalidate(6, code = $$props.code);
    	};

    	$$self.$capture_state = () => ({
    		MultiSelect,
    		username,
    		password,
    		email,
    		user_group,
    		status,
    		result,
    		code,
    		value,
    		handleClick
    	});

    	$$self.$inject_state = $$props => {
    		if ('username' in $$props) $$invalidate(0, username = $$props.username);
    		if ('password' in $$props) $$invalidate(1, password = $$props.password);
    		if ('email' in $$props) $$invalidate(2, email = $$props.email);
    		if ('user_group' in $$props) $$invalidate(3, user_group = $$props.user_group);
    		if ('status' in $$props) $$invalidate(4, status = $$props.status);
    		if ('result' in $$props) $$invalidate(5, result = $$props.result);
    		if ('code' in $$props) $$invalidate(6, code = $$props.code);
    		if ('value' in $$props) value = $$props.value;
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [
    		username,
    		password,
    		email,
    		user_group,
    		status,
    		result,
    		code,
    		handleClick,
    		input0_input_handler,
    		input1_input_handler,
    		input2_input_handler,
    		multiselect_value_binding,
    		select_change_handler
    	];
    }

    class AdminUpdateUser extends SvelteComponentDev {
    	constructor(options) {
    		super(options);

    		init(this, options, instance, create_fragment, safe_not_equal, {
    			username: 0,
    			password: 1,
    			email: 2,
    			user_group: 3,
    			status: 4,
    			result: 5,
    			code: 6
    		});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "AdminUpdateUser",
    			options,
    			id: create_fragment.name
    		});
    	}

    	get username() {
    		throw new Error("<AdminUpdateUser>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set username(value) {
    		throw new Error("<AdminUpdateUser>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get password() {
    		throw new Error("<AdminUpdateUser>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set password(value) {
    		throw new Error("<AdminUpdateUser>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get email() {
    		throw new Error("<AdminUpdateUser>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set email(value) {
    		throw new Error("<AdminUpdateUser>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get user_group() {
    		throw new Error("<AdminUpdateUser>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set user_group(value) {
    		throw new Error("<AdminUpdateUser>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get status() {
    		throw new Error("<AdminUpdateUser>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set status(value) {
    		throw new Error("<AdminUpdateUser>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get result() {
    		throw new Error("<AdminUpdateUser>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set result(value) {
    		throw new Error("<AdminUpdateUser>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	get code() {
    		throw new Error("<AdminUpdateUser>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set code(value) {
    		throw new Error("<AdminUpdateUser>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    const app = new AdminUpdateUser({
    	target: document.body,
    	props: {
    		username: ''
    	}
    });

    return app;

})();
//# sourceMappingURL=bundle.js.map
