import { EventsOn, EventsOff, EventsEmit } from '../../wailsjs/runtime';

export function eventsOn(eventName: string, callback: (data: object) => void) {
  EventsOn(eventName, callback);
}

export function eventsOff(eventName: string) {
  return EventsOff(eventName);
}

export function eventsEmit(eventName: string, data: object) {
  EventsEmit(eventName, data);
}