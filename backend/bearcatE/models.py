from django.db import models

# Create your models here.
class EventModel(models.Model):

    title = models.CharField(max_length = 100, help_text='title of event')
    description = models.TextField(help_text='description about event')

    
    def __str__(self):
        return self.title

